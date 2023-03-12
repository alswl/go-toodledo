package app

import (
	"sync"
	"time"

	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/alswl/go-toodledo/pkg/ui/primarypane"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/alswl/go-toodledo/pkg/ui"

	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/common/logging"
	"github.com/alswl/go-toodledo/pkg/fetchers"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/services"
	uisidebar "github.com/alswl/go-toodledo/pkg/ui/sidebar"
	uistatusbar "github.com/alswl/go-toodledo/pkg/ui/statusbar"
	"github.com/sirupsen/logrus"
)

// const mainModel = "main"

const sidebarModel = "sidebar"
const statusbarModel = "statusbar"
const defaultAutoSyncDuration = 5 * time.Minute
const EnvTTNoFetch = "TT_NO_FETCH"

// TODO move it
var refreshLock sync.Mutex

// const defaultSyncTimeout = 2 * 60 * time.Second.

var _ tea.Model = (*Model)(nil)

type States struct {
	// Tasks is available tasks
	Tasks    []*models.RichTask
	Contexts []*models.Context
	Folders  []*models.Folder
	Goals    []*models.Goal

	// query is current query of task pane
	query *queries.TaskListQuery

	isSidebarVisible bool
}

func NewStates() *States {
	return &States{
		Tasks:            []*models.RichTask{},
		Contexts:         []*models.Context{},
		Folders:          []*models.Folder{},
		Goals:            []*models.Goal{},
		query:            queries.NewTaskListQuery(),
		isSidebarVisible: true,
	}
}

// Model is the main tt app
// it was singleton.
type Model struct {
	ui.Resizable
	ui.Focusable
	ui.Containerized

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

	// TODO ready check
	ready bool
	// TODO remove this, using focus to statusbar
	isInputting bool

	// view
	primaryPane primarypane.Model
	sidebar     uisidebar.Model
	statusBar   uistatusbar.Model
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
		Containerized: *ui.NewContainerized(ui.PrimaryModel, []string{
			ui.PrimaryModel,
			sidebarModel,
		}),
		ready:       false,
		statusBar:   statusBar,
		sidebar:     sidebar,
		isInputting: false,
	}

	primaryPane := primarypane.InitModel(
		*primarypane.NewProperties(
			m.states.query,
			m.states.Tasks,
			m.handleCompleteToggle,
			m.handleTimerToggle,
			m.handleEditTask,
		), m.taskRichSvc,
		m.Width,
		m.Height,
	)
	m.primaryPane = primaryPane

	// init fetcher
	describer := fetchers.NewStatusDescriber(func() error {
		//// TODO using register fun instead of invoke m in ModeNew func
		m.statusBar.Info("fetching...")
		return nil
	}, func() error {
		// TODO using register fun instead of invoke m in ModeNew func
		now := time.Now()
		m.statusBar.Info("fetch done at " + now.Format("15:04:05"))
		return nil
	}, func(err error) error {
		// TODO using register fun instead of invoke m in ModeNew func
		m.statusBar.Error("fetch error: " + err.Error())
		return nil
	})
	interval, err := time.ParseDuration(config.AutoRefresh)
	if err != nil {
		log.WithField("duration", config.AutoRefresh).Error("parse duration error")
		interval = defaultAutoSyncDuration
	}
	fetcher := fetchers.NewSimpleFetcher(log, interval, fetchers.NewToodledoFetchFnPartial(
		log,
		app.FolderExtSvc,
		app.ContextExtSvc,
		app.GoalExtSvc,
		app.TaskExtSvc,
		app.AccountSvc,
	), describer)
	// TODO using register fun instead of invoke m in ModeNew func
	m.fetcher = fetcher

	m.primaryPane.Focus()

	return &m, nil
}
