package ui

import (
	"github.com/alswl/go-toodledo/cmd/tt/ui/styles"
	"github.com/alswl/go-toodledo/cmd/tt/ui/utils"
	"github.com/alswl/go-toodledo/pkg/models"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

type Model struct {
	keys utils.KeyMap
	err  error
	//config
	data []*models.RichTask

	tasksModel   TasksPane
	sidebarModel SidebarPane

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
				m.tasksModel = newM.(TasksPane)
			} else if m.activateModel == "sidebar" {
				newM, _ := m.sidebarModel.Update(msg)
				m.sidebarModel = newM.(SidebarPane)
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

func InitialModel() Model {
	//ts, err := AllTasks()
	// FIXME
	_, err := AllTasksMock()

	if err != nil {
		_ = []*models.RichTask{}
	}

	keys := table.DefaultKeyMap()
	keys.RowDown.SetKeys("j", "down")
	keys.RowUp.SetKeys("k", "up")

	m := Model{
		tasksModel:    InitialTasksPane(),
		sidebarModel:  InitSidebarPane(),
		activateModel: "tasks",
	}
	m.tasksModel.isFocused = true

	return m
}
