package main

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/models"
	tstatus "github.com/alswl/go-toodledo/pkg/models/enums/tasks/status"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	choices    []string         // items on the to-do list
	cursor     int              // which to-do list item our cursor is pointing at
	selected   map[int]struct{} // which to-do items are selected
	tableModel table.Model
}

type item struct {
	models.RichTask
}

func (i item) Title() string       { return i.RichTask.Title }
func (i item) Description() string { return tstatus.StatusValue2Type(i.RichTask.Status).String() }
func (i item) FilterValue() string { return i.RichTask.Title }

func initViper() {
	// Find home directory.
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// Search config in home directory with name ".toodledo" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".toodledo")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		logrus.Debug("config file", viper.ConfigFileUsed())
	}
}

// FIXME using daemon syncer
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

func AllTasksMock() ([]*models.RichTask, error) {
	return []*models.RichTask{
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
			TheGoal:    models.Goal{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "abc"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
		{
			Task:       models.Task{Title: "def"},
			TheContext: models.Context{},
			TheFolder:  models.Folder{},
		},
	}, nil
}

const (
	columnKeyID      = "id"
	columnKeyTitle   = "title"
	columnKeyContext = "context"
	columnKeyStatus  = "status"
)

func initialModel() model {
	ts, err := AllTasks()
	// FIXME
	//ts, err := AllTasksMock()

	if err != nil {
		ts = []*models.RichTask{}
	}

	columns := []table.Column{
		table.NewColumn(columnKeyID, "ID", 15).WithFiltered(true).WithStyle(lipgloss.NewStyle().Faint(true).Foreground(lipgloss.Color("#88f"))),
		table.NewColumn(columnKeyTitle, "Title", 50).WithFiltered(true),
		table.NewColumn(columnKeyContext, "Context", 15),
		table.NewColumn(columnKeyStatus, "Status", 15),
	}

	var rows []table.Row
	for _, t := range ts {
		rows = append(rows, table.NewRow(
			table.RowData{
				columnKeyID:      strconv.Itoa(int(t.ID)),
				columnKeyTitle:   t.Title,
				columnKeyContext: t.TheContext.Name,
				columnKeyStatus:  tstatus.StatusValue2Type(t.Status),
			},
		))
	}

	keys := table.DefaultKeyMap()
	keys.RowDown.SetKeys("j", "down")
	keys.RowUp.SetKeys("k", "up")

	model := model{
		tableModel: table.New(columns).
			WithRows(rows).
			HeaderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)).
			SelectableRows(false).
			Focused(true).
			Filtered(true).
			//Border(customBorder).
			WithKeyMap(keys).
			WithStaticFooter("Footer!").
			WithPageSize(20),
	}

	model.updateFooter()

	return model
}

func (m *model) updateFooter() {
	highlightedRow := m.tableModel.HighlightedRow()

	footerText := fmt.Sprintf(
		"Pg. %d/%d - Currently looking at ID: %s",
		m.tableModel.CurrentPage(),
		m.tableModel.MaxPages(),
		highlightedRow.Data[columnKeyID],
	)

	m.tableModel = m.tableModel.WithStaticFooter(footerText)
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	m.tableModel, cmd = m.tableModel.Update(msg)
	cmds = append(cmds, cmd)

	// We control the footer text, so make sure to update it
	m.updateFooter()

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			cmds = append(cmds, tea.Quit)
		}
	case tea.WindowSizeMsg:
		//top, right, bottom, left := docStyle.GetMargin()
		//m.tableModel.SetSize(msg.Width-left-right, msg.Height-top-bottom)
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	body := strings.Builder{}

	body.WriteString("Press left/right or page up/down to move pages\n")
	body.WriteString("Press space/enter to select a row, q or ctrl+c to quit\n")

	//var selectedIDs []string
	//for _, row := range m.tableModel.SelectedRows() {
	//	// Slightly dangerous type assumption but fine for demo
	//	selectedIDs = append(selectedIDs, row.Data[columnKeyID].(string))
	//}
	//body.WriteString(fmt.Sprintf("SelectedIDs: %s\n", strings.Join(selectedIDs, ", ")))
	body.WriteString(m.tableModel.View())
	body.WriteString("\n")

	return body.String()
}

func main() {
	// TODO
	initViper()

	// TODO full screen
	//p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
