package app

import (
	"fmt"
	"time"

	comsidebar "github.com/alswl/go-toodledo/pkg/ui/sidebar"
	comstatusbar "github.com/alswl/go-toodledo/pkg/ui/statusbar"
	"github.com/alswl/go-toodledo/pkg/ui/taskspane"

	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/common/logging"
	"github.com/alswl/go-toodledo/pkg/fetchers"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/sirupsen/logrus"
)

type States struct {
	width  int
	height int

	// Tasks is available tasks
	Tasks    []*models.RichTask
	Contexts []*models.Context
	Folders  []*models.Folder
	Goals    []*models.Goal
	query    *queries.TaskListQuery
}

type RefreshMsg struct {
	isHardRefresh bool
}

type RefreshTasks struct {
}

// Model is the main tt app
// it was singleton.
type Model struct {
	taskRichSvc   services.TaskRichService
	contextExtSvc services.ContextPersistenceService
	folderExtSvc  services.FolderPersistenceService
	goalExtSvc    services.GoalPersistenceService
	taskExtSvc    services.TaskExtendedService
	taskLocalSvc  services.TaskPersistenceExtService

	// properties
	log logrus.FieldLogger

	// states TODO
	states *States
	err    error
	// focused model: tasks, sidebar, statusbar
	focused string
	// TODO ready check
	ready bool
	// isSidebarOpen bool

	// view
	tasksPanes map[string]*taskspane.Model
	sidebar    comsidebar.Model
	statusBar  comstatusbar.Model
	// TODO help pane
	// help          help.Model
	isInputting bool
	fetcher     fetchers.DaemonFetcher
}

func InitialModel() (*Model, error) {
	// prepare
	log := logging.GetLogger("tt")
	var err error
	app, err := injector.InitTUIApp()
	if err != nil {
		return nil, err
	}
	config, err := injector.InitCLIOption()
	if err != nil {
		return nil, err
	}
	accountSvc := app.AccountSvc
	taskExtSvc := app.TaskExtSvc
	taskLocalSvc := app.TaskExtSvc
	taskRichSvc := app.TaskRichSvc
	contextSvc := app.ContextExtSvc
	folderSvc := app.FolderExtSvc
	goalSvc := app.GoalExtSvc

	_, _, err = accountSvc.CachedMe()
	if err != nil {
		return nil, err
	}

	states := &States{
		Tasks:    []*models.RichTask{},
		Contexts: []*models.Context{},
		Folders:  []*models.Folder{},
		Goals:    []*models.Goal{},
		query:    &queries.TaskListQuery{},
	}

	// status bar
	statusBar := comstatusbar.NewDefault()
	statusBar.SetMode("tasks")
	statusBar.SetInfo1(fmt.Sprintf("./%d", len(states.Tasks)))
	statusBar.SetInfo2("HELP(h)")

	// task pane
	sidebar := comsidebar.InitModel(comsidebar.Properties{})

	// main app
	m := Model{
		log:           log,
		taskRichSvc:   taskRichSvc,
		contextExtSvc: contextSvc,
		folderExtSvc:  folderSvc,
		goalExtSvc:    goalSvc,
		taskExtSvc:    taskExtSvc,
		taskLocalSvc:  taskLocalSvc,
		states:        states,
		err:           nil,
		focused:       "tasks",
		ready:         false,
		statusBar:     statusBar,
		sidebar:       sidebar,
		isInputting:   false,
		tasksPanes:    map[string]*taskspane.Model{},
	}

	m.sidebar.RegisterItemChange(m.OnItemChange)

	// init fetcher
	describer := fetchers.NewStatusDescriber(func() error {
		// TODO using register fun instead of invoke m in New func
		m.statusBar.SetStatus("fetching...")
		return nil
	}, func() error {
		// TODO using register fun instead of invoke m in New func
		m.statusBar.SetStatus("fetching done")
		return nil
	}, func(err error) error {
		// TODO using register fun instead of invoke m in New func
		m.statusBar.SetStatus("fetching error: " + err.Error())
		return nil
	})
	duration, err := time.ParseDuration(config.AutoRefresh)
	if err != nil {
		log.WithField("duration", config.AutoRefresh).Error("parse duration error")
		duration = defaultAutoSyncDuration
	}
	fetcher := fetchers.NewSimpleFetcher(log, duration, fetchers.NewToodledoFetchFnPartial(
		log,
		app.FolderExtSvc,
		app.ContextExtSvc,
		app.GoalExtSvc,
		app.TaskExtSvc,
		app.AccountSvc,
	), describer)
	// TODO using register fun instead of invoke m in New func
	m.fetcher = fetcher

	m.getOrCreateTaskPaneByQuery().Focus()

	return &m, nil
}
