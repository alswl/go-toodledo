package ui

import (
	"github.com/alswl/go-toodledo/cmd/tt/ui/utils"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

type Model struct {
	keys utils.KeyMap
	err  error
	//config
	data []*models.RichTask

	tasksModel TasksPane
	sidebar    Sidebar

	help          help.Model
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
		default:
			// bubble event to sub component
			m.tasksModel, _ = m.tasksModel.Update(msg)
		}
	}
	return m, nil
}

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	paddedContentStyle := lipgloss.NewStyle().
		Padding(0, mainContentPadding)

	return paddedContentStyle.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			m.sidebar.View(),
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
		tasksModel: InitialTasksPane(),
		sidebar:    InitSidebarPane(),
	}

	return m
}
