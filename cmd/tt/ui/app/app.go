package app

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/cmd/tt/ui"
	"github.com/alswl/go-toodledo/cmd/tt/ui/components/sidebar"
	"github.com/alswl/go-toodledo/cmd/tt/ui/components/taskspane"
	"github.com/alswl/go-toodledo/cmd/tt/ui/styles"
	"github.com/alswl/go-toodledo/cmd/tt/ui/utils"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/sirupsen/logrus"
	"os"
)

type States struct {
	// Tasks is available tasks
	Tasks []*models.RichTask
}

type Model struct {
	keys utils.KeyMap
	err  error
	//config
	data   []*models.RichTask
	states *States

	tasksModel   taskspane.Model
	sidebarModel sidebar.Model

	//activateModel tea.Model
	activateModel string

	// TODO help pane
	//help          help.Model
	// TODO ready check
	ready         bool
	isSidebarOpen bool
	width         int
	//filterWindow  FilterFormModel
	//tabs            tabs.Model
	//context         context.ProgramContext
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// TODO keymap
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			// TODO refactor, switch with mod
			if m.activateModel == "tasks" {
				m.activateModel = "sidebar"
				m.sidebarModel.Focus()
				m.tasksModel.Blur()
			} else if m.activateModel == "sidebar" {
				m.activateModel = "tasks"
				m.tasksModel.Focus()
				m.sidebarModel.Blur()
			}
			return m, nil
		default:
			// bubble event to sub component
			if m.activateModel == "tasks" {
				newM, _ := m.tasksModel.Update(msg)
				m.tasksModel = newM.(taskspane.Model)
			} else if m.activateModel == "sidebar" {
				newM, _ := m.sidebarModel.Update(msg)
				m.sidebarModel = newM.(sidebar.Model)
			}
		}
	case tea.WindowSizeMsg:
		sideBarWidth := msg.Width / 12 * 3
		m.sidebarModel.Resize(sideBarWidth, msg.Height)
		m.tasksModel.Resize(msg.Width-sideBarWidth, msg.Height)
	}
	return m, nil
}

func (m Model) View() string {
	// TODO error handling
	if m.err != nil {
		return m.err.Error()
	}

	return styles.MainPaneStyle.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.sidebarModel.View(),
			m.tasksModel.View(),
		),
	)
}

//func (m Model) renderTableHeader() string {
//	return headerStyle.
//		PaddingLeft(mainContentPadding).
//		PaddingRight(mainContentPadding).
//		Width(m.mainViewport.model.Width).
//		MaxWidth(m.mainViewport.model.Width).
//		Render(
//			lipgloss.JoinHorizontal(
//				lipgloss.Left,
//				//updatedAtCell,
//				//reviewCell,
//				//prRepoCell,
//				//prTitleCell,
//				//prAuthorCell,
//				//mergeableCell,
//				//ciCell,
//				//linesCell,
//			),
//		)
//}

//func (m Model) renderCurrentSection() string {
//	return lipgloss.NewStyle().
//		PaddingLeft(mainContentPadding).
//		PaddingRight(mainContentPadding).
//		MaxWidth(m.mainViewport.model.Width).
//		Render(m.RenderMainViewPort())
//}

// FIXME using daemon syncer
// nolint:deadcode
func AllTasks() ([]*models.RichTask, error) {
	_, err := injector.InitApp()
	if err != nil {
		logrus.Fatal("login required, using `toodledo auth login` to login.")
		return nil, err
	}
	svc, err := injector.InitTaskCachedService()
	if err != nil {
		logrus.WithError(err).Fatal("failed to init task service")
		return nil, err
	}
	syncer, err := injector.InitSyncer()
	if err != nil {
		logrus.WithError(err).Fatal("init syncer failed")
		return nil, err
	}
	taskRichSvc, err := injector.InitTaskRichService()
	if err != nil {
		logrus.WithError(err).Fatal("init task rich service failed")
		return nil, err
	}
	err = syncer.SyncOnce()
	if err != nil {
		logrus.WithError(err).Fatal("sync failed")
		return nil, err
	}
	tasks, err := svc.ListAllByQuery(&queries.TaskListQuery{})
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	rts, err := taskRichSvc.RichThem(tasks)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return rts, nil
}

func InitialModel() Model {
	var err error
	var tasks []*models.RichTask
	if os.Getenv("MOCK") == "true" {
		tasks, err = ui.AllTasksMock()
	} else {
		tasks, err = AllTasks()
	}
	if err != nil {
		// FIXME
		panic(err)
	}

	// FIXME move tables km to model
	keys := table.DefaultKeyMap()
	keys.RowDown.SetKeys("j", "down")
	keys.RowUp.SetKeys("k", "up")

	m := Model{
		tasksModel:   taskspane.InitModel(tasks),
		sidebarModel: sidebar.InitModel(),
	}
	// FIXME focus as an method
	m.activateModel = "tasks"
	m.tasksModel.Focus()

	return m
}
