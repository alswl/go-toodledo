package main

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/models"
	tstatus "github.com/alswl/go-toodledo/pkg/models/enums/tasks/status"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	tealist "github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
	list     tealist.Model
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

func initialModel() model {
	ts, err := AllTasks()
	if err != nil {
		ts = []*models.RichTask{}
	}

	var items []tealist.Item
	for _, t := range ts {
		items = append(items, item{*t})
	}

	m := model{list: tealist.New(items, tealist.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "My Tasks"
	return m
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// main window
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// These keys should exit the program.
		case "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		top, right, bottom, left := docStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom)
	}

	// list
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func main() {
	initViper()
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
