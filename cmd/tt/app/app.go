package app

import (
	"context"
	"sync"
	"time"

	"github.com/alswl/go-toodledo/pkg/models/queries"

	"github.com/alswl/go-toodledo/pkg/ui/sidebar"
	"gopkg.in/yaml.v3"

	"github.com/alswl/go-toodledo/pkg/models"

	tea "github.com/charmbracelet/bubbletea"
)

// const defaultSyncTimeout = 2 * 60 * time.Second.
const defaultAutoSyncDuration = 5 * time.Minute

var (
	switchAllowedPanes = []string{
		"tasks",
		"sidebar",
	}
)

// TODO move it
var refreshLock sync.Mutex

func (m *Model) Init() tea.Cmd {
	var cmds []tea.Cmd

	cmds = append(cmds, m.ReloadDependencies())

	// using default first now
	// if len(m.states.Contexts) > 0 {
	//	m.states.query.ContextID = m.states.Contexts[0].ID
	//}
	cmds = append(cmds, m.ReloadTasks())

	// tasks
	// m.statusBar.SetStatus(fmt.Sprintf("INFO: tasks: %d", len(tasks)))

	// states init
	cmds = append(cmds, func() tea.Msg {
		ierr := m.contextExtSvc.Sync()
		if ierr != nil {
			m.err = ierr
			return nil
		}
		ierr = m.contextExtSvc.Sync()
		if ierr != nil {
			m.err = ierr
			return nil
		}
		ierr = m.folderExtSvc.Sync()
		if ierr != nil {
			m.err = ierr
			return nil
		}
		ierr = m.goalExtSvc.Sync()
		if ierr != nil {
			m.err = ierr
			return nil
		}

		return models.RefreshPropertiesMsg{}
	})

	// daemon fetcher sstart
	// XXX using tae.Every
	cmds = append(cmds, func() tea.Msg {
		m.fetcher.Start(context.Background())
		return nil
	})

	// refresh at start
	cmds = append(cmds, func() tea.Msg {
		return models.FetchTasksMsg{IsHardRefresh: false}
	})

	// statusbar Init // TODO should I call it manually?
	cmds = append(cmds, m.statusBar.Init())

	// update last sidebar setting
	cmds = append(cmds, func() tea.Msg {
		bs, err := m.settingSvc.Find(sidebarStatesKey)
		if err != nil {
			m.log.WithError(err).Error("get sidebar states failed")
			return nil
		}
		states := sidebar.NewStates()
		err = yaml.Unmarshal([]byte(bs), &states)
		if err != nil {
			m.log.WithError(err).Error("unmarshal sidebar states failed")
			return nil
		}
		var cmd tea.Cmd
		m.sidebar, cmd = m.sidebar.UpdateTyped(states)
		return cmd
	})
	query := queries.TaskListQuery{}
	lastQuery, err := m.settingSvc.Find(lastQueryKey)
	if err == nil {
		_ = yaml.Unmarshal([]byte(lastQuery), &query)
	}
	m.states.query = &query

	return tea.Batch(cmds...)
}
