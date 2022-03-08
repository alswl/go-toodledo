package ui

import (
	"github.com/alswl/go-toodledo/pkg/ui/utils"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type Model struct {
	keys utils.KeyMap
	err  error
	//config          *config.Config
	//data            *[]Section
	mainViewport    MainViewport
	sidebarViewport viewport.Model
	cursor          cursor
	help            help.Model
	ready           bool
	isSidebarOpen   bool
	width           int
	filterWindow    FilterFormModel
	//tabs            tabs.Model
	//context         context.ProgramContext
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// TODO keymap
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

	}
	return m, nil
}

func (m *Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	paddedContentStyle := lipgloss.NewStyle().
		Padding(0, mainContentPadding)

	s := strings.Builder{}
	s.WriteString(m.filterWindow.View())
	s.WriteString("\n")
	table := paddedContentStyle.Render(lipgloss.JoinVertical(lipgloss.Top, m.renderTableHeader(), m.renderCurrentSection()))
	s.WriteString(table)
	s.WriteString("\n")
	// TODO
	//s.WriteString(m.renderHelp())
	return s.String()
}

func (m *Model) renderTableHeader() string {

	return headerStyle.
		PaddingLeft(mainContentPadding).
		PaddingRight(mainContentPadding).
		Width(m.mainViewport.model.Width).
		MaxWidth(m.mainViewport.model.Width).
		Render(
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				//updatedAtCell,
				//reviewCell,
				//prRepoCell,
				//prTitleCell,
				//prAuthorCell,
				//mergeableCell,
				//ciCell,
				//linesCell,
			),
		)
}

func (m *Model) renderCurrentSection() string {
	return lipgloss.NewStyle().
		PaddingLeft(mainContentPadding).
		PaddingRight(mainContentPadding).
		MaxWidth(m.mainViewport.model.Width).
		Render(m.RenderMainViewPort())
}

func (m *Model) RenderMainViewPort() string {
	return "\n"
}

type MainViewport struct {
	model viewport.Model
}

type cursor struct {
	currID int
}
