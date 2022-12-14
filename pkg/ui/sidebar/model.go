package sidebar

import (
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/common/logging"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/constants"
	"github.com/alswl/go-toodledo/pkg/ui"
	"github.com/alswl/go-toodledo/pkg/ui/common"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sirupsen/logrus"
)

type Properties struct {
}

var defaultTabs = []string{
	constants.Contexts,
	constants.Folders,
	constants.Goals,
	// "Priority",
	// "Tags",
	// "Search",
}

type Model struct {
	ui.Focusable
	ui.Resizable

	// props
	log        logrus.FieldLogger
	properties Properties

	// states
	isCollapsed     bool
	currentTabIndex int
	currentTab      string
	Contexts        []models.Context
	Folders         []models.Folder
	Goals           []models.Goal

	// view
	// list has states(selected)
	// TODO using wrapped list
	contextList list.Model
	folderList  list.Model
	goalList    list.Model

	// handler
	onItemChangeSubscribers []ItemChangeSubscriber
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	// TODO move to styles
	tabName := defaultTabs[m.currentTabIndex]
	tabRender := styles.EmptyStyle.Render("<" + tabName + ">")

	m.Viewport.SetContent(
		lipgloss.JoinVertical(
			lipgloss.Left,
			tabRender,
			styles.EmptyStyle.Render(m.getVisibleList().View()),
		),
	)
	style := styles.PaneStyle.Copy()
	if m.IsFocused() {
		style = styles.FocusedPaneStyle.Copy()
	}
	// return style.
	//	Width(m.Viewport.Width).
	//	Height(m.Viewport.Height).
	//	Render(wrap.String(
	//		wordwrap.String(m.Viewport.View(), m.Viewport.Width), m.Viewport.Width),
	//	)
	return style.Width(m.Viewport.Width).Render(m.Viewport.View())
}

func (m Model) UpdateTyped(msg tea.Msg) (Model, tea.Cmd) {
	newM, cmd := m.Update(msg)
	return newM.(Model), cmd
}

func (m *Model) updateTab(step int) {
	if step == 0 {
		return
	}

	newIndex := (m.currentTabIndex + step + len(defaultTabs)) % len(defaultTabs)
	m.currentTabIndex = newIndex
}

func (m *Model) getVisibleList() *list.Model {
	tab := defaultTabs[m.currentTabIndex]
	var l *list.Model

	switch tab {
	case constants.Contexts:
		l = &m.contextList
	case constants.Folders:
		l = &m.folderList
	case constants.Goals:
		l = &m.goalList
	default:
		panic("unknown tab")
	}
	return l
}

func (m *Model) updateVisibleList(msg tea.Msg) tea.Cmd {
	tab := defaultTabs[m.currentTabIndex]
	var cmd tea.Cmd

	switch tab {
	case constants.Contexts:
		m.contextList, cmd = m.contextList.Update(msg)
	case constants.Folders:
		m.folderList, cmd = m.folderList.Update(msg)
	case constants.Goals:
		m.goalList, cmd = m.goalList.Update(msg)
	}
	return cmd
}

func (m *Model) Resize(width, height int) {
	m.Resizable.Resize(width, height, styles.PaneStyle.GetBorderStyle())
}

func (m *Model) RegisterItemChange(onItemChange ItemChangeSubscriber) {
	m.onItemChangeSubscribers = append(m.onItemChangeSubscribers, onItemChange)
}

func InitModel(p Properties) Model {
	m := Model{
		log:                     logging.GetLogger("tt"),
		properties:              p,
		isCollapsed:             false,
		currentTabIndex:         0,
		onItemChangeSubscribers: []ItemChangeSubscriber{},
		contextList:             common.NewSimpleList(),
		folderList:              common.NewSimpleList(),
		goalList:                common.NewSimpleList(),
	}
	// if len(m.list.Items()) > 0 {
	//	m.list.Select(0)
	//}
	m.Blur()
	return m
}
