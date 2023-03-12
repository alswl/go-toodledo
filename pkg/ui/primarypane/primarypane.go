package primarypane

import (
	"github.com/alswl/go-toodledo/cmd/tt/styles"
	"github.com/alswl/go-toodledo/pkg/common/logging"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/alswl/go-toodledo/pkg/ui"
	uidetail "github.com/alswl/go-toodledo/pkg/ui/detail"
	"github.com/alswl/go-toodledo/pkg/ui/taskstablepane"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sirupsen/logrus"
)

type Properties struct {
	Query *queries.TaskListQuery

	Tasks []*models.RichTask

	HandleCompleteToggleFn func(id int64) tea.Cmd
	HandleTimerToggleFn    func(id int64) tea.Cmd
	// TODO self handle
	HandleEditTaskFn func(id int64) tea.Cmd
}

func NewProperties(
	query *queries.TaskListQuery,
	tasks []*models.RichTask,
	handleCompleteToggleFn func(id int64) tea.Cmd,
	handleTimerToggleFn func(id int64) tea.Cmd,
	handleEditTaskFn func(id int64) tea.Cmd,
) *Properties {
	return &Properties{
		Query:                  query,
		Tasks:                  tasks,
		HandleCompleteToggleFn: handleCompleteToggleFn,
		HandleTimerToggleFn:    handleTimerToggleFn,
		HandleEditTaskFn:       handleEditTaskFn,
	}
}

type Model struct {
	ui.Focusable
	ui.Containerized
	ui.Resizable

	taskRichSvc services.TaskRichService

	// TODO ptr or not
	tasksPanes map[string]*taskstablepane.Model
	taskDetail *uidetail.Model

	// TODO ptr or not
	props Properties
	log   logrus.FieldLogger
}

func InitModel(props Properties, taskRichSvc services.TaskRichService, width, height int) Model {
	m := Model{
		// XXX test it, q
		Containerized: *ui.NewContainerized(ui.TasksTableModel, []string{ui.TasksTableModel, ui.DetailModel}),

		taskRichSvc: taskRichSvc,
		tasksPanes:  map[string]*taskstablepane.Model{},
		taskDetail:  nil,

		props: props,
		// TODO using specific logger
		log: logging.GetLoggerOrDefault("tt"),
	}
	m.Width = width
	m.Height = height
	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Resize(width, height int) {
	if width <= 0 || height <= 0 {
		return
	}

	m.Resizable.Resize(width, height, styles.NoStyle.Copy().GetBorderStyle())

	if m.taskDetail != nil {
		m.taskDetail.Resize(width, height)
	}

	for _, p := range m.tasksPanes {
		p.Resize(width, height-1)
	}
}

func (m Model) View() string {
	if m.taskDetail != nil {
		return m.taskDetail.View()
	}

	taskPane := m.GetOrCreateTaskPaneByQuery()

	style := styles.PaneStyle.Copy()
	if m.IsFocused() {
		style = styles.FocusedPaneStyle.Copy()
	}
	return style.Render(taskPane.View())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	focused := m.Focused()
	if focused == ui.PrimaryModel {
		// task table is default
		focused = ui.TasksTableModel
	}

	// nolint
	switch msgTyped := msg.(type) {
	case tea.KeyMsg:
		switch focused {
		case ui.DetailModel:
			var newM uidetail.Model
			newM, cmd = m.taskDetail.UpdateTyped(msg)
			m.taskDetail = &newM

		case ui.TasksTableModel:
			// table
			cmd = m.UpdateTaskPane(msg)
		}

	case models.ReturnMsg:
		// return from subcomponent
		m.taskDetail = nil
		m.FocusChild(ui.TasksTableModel)

	case models.RichTask:
		// TODO delete?
		m.taskDetail = uidetail.InitModel(msgTyped, m.Width, m.Height)
		m.taskDetail.Resize(m.Width, m.Height)

	case tea.WindowSizeMsg:
		m.Resize(msgTyped.Width, msgTyped.Height)
		// viewport must set content in every sizing
		// example, https://github.com/charmbracelet/bubbletea/blob/master/examples/pager/main.go#L74
		m.Viewport.SetContent(m.View())
	}

	// m.Viewport, cmd = m.Viewport.Update(msg)
	return m, cmd
}

func (m Model) UpdateTyped(msg tea.Msg) (Model, tea.Cmd) {
	newM, cmd := m.Update(msg)
	return newM.(Model), cmd
}

// GetOrCreateTaskPaneByQuery returns the task pane by query.
// app holds a map of task panes, and the key is the query string.
func (m *Model) GetOrCreateTaskPaneByQuery() *taskstablepane.Model {
	key := ""
	if m.props.Query != nil {
		key = m.props.Query.UniqString()
	}
	if p, ok := m.tasksPanes[key]; ok {
		return p
	}
	newP := taskstablepane.InitModel(m.props.Tasks, m.Width, m.Height)
	m.tasksPanes[key] = &newP
	return &newP
}

func (m Model) GetPane() taskstablepane.Model {
	return *m.GetOrCreateTaskPaneByQuery()
}

func (m *Model) UpdateTaskPane(msg tea.Msg) tea.Cmd {
	paneKey := ""
	if m.props.Query != nil {
		paneKey = m.props.Query.UniqString()
	}
	// pane := m.GetOrCreateTaskPaneByQuery()

	pane := m.GetPane()
	id, err := pane.Selected()
	if err != nil {
		m.log.WithError(err).Warn("get selected task")
	}

	var cmd tea.Cmd
	switch msgType := msg.(type) {
	case tea.KeyMsg:
		switch msgType.String() {
		case "x":
			cmd = m.props.HandleCompleteToggleFn(id)
		case "enter":
			cmd = m.props.HandleTimerToggleFn(id)
		case "o":
			// open task detail pane
			cmd = m.handleOpenTaskDetail(id)
			m.FocusChild(ui.DetailModel)
		case "e":
			cmd = m.props.HandleEditTaskFn(id)
		default:
			var newPane taskstablepane.Model
			newPane, cmd = pane.UpdateTyped(msg)
			m.tasksPanes[paneKey] = &newPane
		}
	default:
		var newPane taskstablepane.Model
		newPane, cmd = pane.UpdateTyped(msg)
		m.tasksPanes[paneKey] = &newPane
	}

	return cmd
}

func (m *Model) handleOpenTaskDetail(id int64) tea.Cmd {
	// m.states.taskDetailID = id

	taskRich, err := m.taskRichSvc.Find(id)
	if err != nil {
		m.log.WithField("id", id).WithError(err).Warn("get task by id")
		return nil
	}

	// newM, cmd := m.UpdateTyped(*taskRich)
	m.taskDetail = uidetail.InitModel(*taskRich, m.Width, m.Height)
	m.taskDetail.Focus()
	return tea.Batch(func() tea.Msg {
		return models.RefreshUIMsg{}
	})
}
