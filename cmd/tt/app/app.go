package app

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/cmd/tt/components/sidebar"
	"github.com/alswl/go-toodledo/cmd/tt/components/taskspane"
	"github.com/alswl/go-toodledo/cmd/tt/mockdata"
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/cmd/tt/utils"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/knipferrc/teacup/statusbar"
	"github.com/sirupsen/logrus"
	"os"
)

type States struct {
	// Tasks is available tasks
	Tasks []*models.RichTask
	// XXX
	Filter string
}

type Model struct {
	keys utils.KeyMap
	err  error
	//config
	data   []*models.RichTask
	states *States

	tasksPane taskspane.Model
	sidebar   sidebar.Model
	statusbar statusbar.Bubble

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
	isInputting bool
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
				m.sidebar.Focus()
				m.tasksPane.Blur()
			} else if m.activateModel == "sidebar" {
				m.activateModel = "tasks"
				m.tasksPane.Focus()
				m.sidebar.Blur()
			}
			return m, nil
		default:
			// bubble event to sub component
			if m.activateModel == "tasks" {
				newM, _ := m.tasksPane.Update(msg)
				m.tasksPane = newM.(taskspane.Model)
			} else if m.activateModel == "sidebar" {
				newM, _ := m.sidebar.Update(msg)
				m.sidebar = newM.(sidebar.Model)
			}
		}
	case tea.WindowSizeMsg:
		sideBarWidth := msg.Width / 12 * 3
		m.sidebar.Resize(sideBarWidth, msg.Height-1)
		m.tasksPane.Resize(msg.Width-sideBarWidth, msg.Height-1)
	}
	return m, nil
}

func (m Model) View() string {
	// TODO error handling
	if m.err != nil {
		return m.err.Error()
	}

	return styles.MainPaneStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				m.sidebar.View(),
				m.tasksPane.View(),
			),
			m.statusbar.View(),
		),
	)
}

func (m Model) GetStates() States {
	return *m.states
}

//func (m *Model) SetInputting(is bool) {
//	m.isInputting = is
//}

//func (m Model) IsInputting() bool {
//	return m.isInputting
//}

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
		tasks, err = mockdata.AllTasksMock()
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

	statusbar := statusbar.New(
		statusbar.ColorConfig{
			Background: lipgloss.AdaptiveColor{Light: styles.Colors.White},
			Foreground: lipgloss.AdaptiveColor{Light: styles.Colors.Pink},
		},
		statusbar.ColorConfig{
			Background: lipgloss.AdaptiveColor{Light: styles.Colors.White},
			Foreground: lipgloss.AdaptiveColor{Light: styles.Colors.Pink},
		},
		statusbar.ColorConfig{
			Background: lipgloss.AdaptiveColor{Light: styles.Colors.White},
			Foreground: lipgloss.AdaptiveColor{Light: styles.Colors.Pink},
		},
		statusbar.ColorConfig{
			Background: lipgloss.AdaptiveColor{Light: styles.Colors.White},
			Foreground: lipgloss.AdaptiveColor{Light: styles.Colors.Pink},
		},
	)
	statusbar.FirstColumn = "tasks"
	statusbar.SecondColumn = "search"
	statusbar.ThirdColumn = "1/999"
	statusbar.FourthColumn = "status"

	m := Model{
		tasksPane: taskspane.InitModel(tasks),
		sidebar:   sidebar.InitModel(),
		statusbar: statusbar,
	}
	// FIXME focus as an method
	m.activateModel = "tasks"
	m.tasksPane.Focus()

	return m
}
