package app

import (
	"context"
	"os"

	"github.com/alswl/go-toodledo/pkg/models/queries"

	"github.com/alswl/go-toodledo/pkg/ui/sidebar"
	"gopkg.in/yaml.v3"

	"github.com/alswl/go-toodledo/pkg/models"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Init() tea.Cmd {
	var cmds []tea.Cmd

	cmds = append(cmds, m.ReloadDependencies())

	query := queries.TaskListQuery{}
	lastQuery, err := m.settingSvc.Find(lastQueryKey)
	if err == nil {
		_ = yaml.Unmarshal([]byte(lastQuery), &query)
	}
	m.states.query = &query
	cmds = append(cmds, m.LoadTasks())

	// states init
	cmds = append(cmds, func() tea.Msg {
		if os.Getenv(EnvTTNoFetch) != "" {
			return nil
		}
		ierr := m.contextExtSvc.Sync()
		if ierr != nil {
			return models.StatusMsg{Message: ierr.Error()}
		}
		ierr = m.contextExtSvc.Sync()
		if ierr != nil {
			return models.StatusMsg{Message: ierr.Error()}
		}
		ierr = m.folderExtSvc.Sync()
		if ierr != nil {
			return models.StatusMsg{Message: ierr.Error()}
		}
		ierr = m.goalExtSvc.Sync()
		if ierr != nil {
			return models.StatusMsg{Message: ierr.Error()}
		}

		return models.RefreshPropertiesMsg{}
	})

	// daemon fetcher start
	// XXX using tae.Every
	cmds = append(cmds, func() tea.Msg {
		if os.Getenv(EnvTTNoFetch) != "" {
			return nil
		}
		m.fetcher.Start(context.Background())
		return nil
	})

	// refresh at start
	cmds = append(cmds, func() tea.Msg {
		if os.Getenv(EnvTTNoFetch) != "" {
			return nil
		}
		return models.FetchTasksMsg{IsHardRefresh: false}
	})

	// statusbar Init // TODO should I call it manually?
	cmds = append(cmds, m.statusBar.Init())

	// update last sidebar setting
	cmds = append(cmds, func() tea.Msg {
		bs, ierr := m.settingSvc.Find(sidebarStatesKey)
		if ierr != nil {
			m.log.WithError(ierr).Error("get sidebar states failed")
			return nil
		}
		states := sidebar.NewStates()
		ierr = yaml.Unmarshal([]byte(bs), &states)
		if ierr != nil {
			m.log.WithError(ierr).Error("unmarshal sidebar states failed")
			return nil
		}
		var cmd tea.Cmd
		m.sidebar, cmd = m.sidebar.UpdateTyped(states)
		return cmd
	})

	return tea.Batch(cmds...)
}
