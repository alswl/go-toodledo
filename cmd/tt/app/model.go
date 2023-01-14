package app

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/alswl/go-toodledo/pkg/ui"

	uidetail "github.com/alswl/go-toodledo/pkg/ui/detail"
	uisidebar "github.com/alswl/go-toodledo/pkg/ui/sidebar"
	uistatusbar "github.com/alswl/go-toodledo/pkg/ui/statusbar"
	"github.com/alswl/go-toodledo/pkg/ui/taskspane"

	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/common/logging"
	"github.com/alswl/go-toodledo/pkg/fetchers"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/sirupsen/logrus"
)

const mainModel = "main"

// taskModels is a list of task models
// tasksModels is not focusable, it was controlled by main
// const tasksModels = "tasksModels".
const sidebarModel = "sidebar"
const statusbarModel = "statusbar"
const detailModel = "detail"

var _ tea.Model = (*Model)(nil)

type States struct {
	// Tasks is available tasks
	Tasks    []*models.RichTask
	Contexts []*models.Context
	Folders  []*models.Folder
	Goals    []*models.Goal

	// query is current query of task pane
	query *queries.TaskListQuery

	// taskDetailID is current task id of detail pane
	taskDetailID int64
}

func NewStates() *States {
	return &States{
		Tasks:        []*models.RichTask{},
		Contexts:     []*models.Context{},
		Folders:      []*models.Folder{},
		Goals:        []*models.Goal{},
		query:        queries.NewTaskListQuery(),
		taskDetailID: 0,
	}
}

// Model is the main tt app
// it was singleton.
type Model struct {
	ui.Resizable
	ui.Focusable

	taskRichSvc   services.TaskRichService
	contextExtSvc services.ContextPersistenceService
	folderExtSvc  services.FolderPersistenceService
	goalExtSvc    services.GoalPersistenceService
	taskExtSvc    services.TaskExtendedService
	taskLocalSvc  services.TaskPersistenceExtService
	settingSvc    services.SettingService
	fetcher       fetchers.DaemonFetcher

	// properties
	log logrus.FieldLogger

	// states TODO
	states *States
	// err    error
	// focused model: main, tasks, sidebar, statusbar
	// focused *ui.Focusable
	focused string

	// TODO ready check
	ready bool
	// TODO remove this, using focus to statusbar
	isInputting bool

	// view
	// TODO ptr or not
	tasksPanes map[string]*taskspane.Model
	sidebar    uisidebar.Model
	statusBar  uistatusbar.Model
	taskDetail uidetail.Model
	// TODO help pane
	// help          help.Model
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
	_, _, err = app.AccountSvc.CachedMe()
	if err != nil {
		return nil, err
	}

	states := NewStates()

	// status bar
	statusBar := uistatusbar.NewDefault()
	statusBar.SetMode(uistatusbar.ModeDefault)
	statusBar.SetInfo1(fmt.Sprintf("./%d", len(states.Tasks)))
	// statusBar.SetInfo2("HELP(h)")

	// task pane
	// XXX sidebar should read query and load specific menu
	sidebar := uisidebar.InitModel(uisidebar.Properties{})

	// main app
	m := Model{
		log:           log,
		taskRichSvc:   app.TaskRichSvc,
		contextExtSvc: app.ContextExtSvc,
		folderExtSvc:  app.FolderExtSvc,
		goalExtSvc:    app.GoalExtSvc,
		taskExtSvc:    app.TaskExtSvc,
		taskLocalSvc:  app.TaskExtSvc,
		settingSvc:    app.SettingSvc,
		states:        states,
		focused:       mainModel,
		ready:         false,
		statusBar:     statusBar,
		sidebar:       sidebar,
		isInputting:   false,
		tasksPanes:    map[string]*taskspane.Model{},
	}

	m.sidebar.RegisterItemChange(m.OnItemChange)

	// init fetcher
	describer := fetchers.NewStatusDescriber(func() error {
		//// TODO using register fun instead of invoke m in ModeNew func
		m.statusBar.SetMessage("fetching...")
		return nil
	}, func() error {
		// TODO using register fun instead of invoke m in ModeNew func
		m.statusBar.SetMessage("fetch done")
		return nil
	}, func(err error) error {
		// TODO using register fun instead of invoke m in ModeNew func
		m.statusBar.SetMessage("fetch error: " + err.Error())
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
	// TODO using register fun instead of invoke m in ModeNew func
	m.fetcher = fetcher

	m.getOrCreateTaskPaneByQuery().Focus()

	return &m, nil
}
