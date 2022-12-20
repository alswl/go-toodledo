package app

import (
	"context"
	"sync"
	"time"

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

	cmds = append(cmds, m.ReloadProperties())
	// using default first now
	// FIXME using last saved view
	m.states.query.ContextID = m.states.Contexts[0].ID
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

	return tea.Batch(cmds...)
}
