package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/utils"
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

	// states init
	cmds = append(cmds, func() tea.Msg {
		err := m.contextExtSvc.Sync()
		if err != nil {
			m.err = err
			return nil
		}
		err = m.contextExtSvc.Sync()
		if err != nil {
			m.err = err
			return nil
		}
		cs, err := m.contextExtSvc.ListAll()
		if err != nil {
			m.err = err
			return nil
		}
		err = m.folderExtSvc.Sync()
		if err != nil {
			m.err = err
			return nil
		}
		fs, err := m.folderExtSvc.ListAll()
		if err != nil {
			m.err = err
			return nil
		}
		err = m.goalExtSvc.Sync()
		if err != nil {
			m.err = err
			return nil
		}
		gs, err := m.goalExtSvc.ListAll()
		if err != nil {
			m.err = err
		}

		// Contexts are first tab in sidebar
		m.states.Contexts = cs
		m.states.Contexts = append([]*models.Context{{
			ID:   0,
			Name: "All",
		}}, cs...)
		m.states.Contexts = append(m.states.Contexts, &models.Context{
			ID:   -1,
			Name: "None",
		})
		m.sidebar, _ = m.sidebar.UpdateTyped(utils.UnwrapListPointer(m.states.Contexts))
		// using default first now
		m.states.query.ContextID = m.states.Contexts[0].ID

		// folders
		m.states.Folders = fs
		m.states.Folders = append([]*models.Folder{{
			ID:   0,
			Name: "All",
		}}, fs...)
		m.states.Folders = append(m.states.Folders, &models.Folder{
			ID:   -1,
			Name: "None",
		})
		m.sidebar, _ = m.sidebar.UpdateTyped(utils.UnwrapListPointer(m.states.Folders))

		// goals
		m.states.Goals = gs
		m.states.Goals = append([]*models.Goal{{
			ID:   0,
			Name: "All",
		}}, gs...)
		m.states.Goals = append(m.states.Goals, &models.Goal{
			ID:   -1,
			Name: "None",
		})
		m.sidebar, _ = m.sidebar.UpdateTyped(utils.UnwrapListPointer(m.states.Goals))

		// TODO using last selected menu

		// tasks
		tasks, err := m.taskExtSvc.ListAllByQuery(m.states.query)
		if err != nil {
			m.statusBar.SetStatus("ERROR: " + err.Error())
		}
		rts, _ := m.taskRichSvc.RichThem(tasks)
		m.states.Tasks = rts

		cmds = append(cmds, m.updateTaskPane(rts))
		m.statusBar.SetStatus(fmt.Sprintf("INFO: tasks: %d", len(tasks)))

		return nil
	})

	// daemon fetcher sstart
	cmds = append(cmds, func() tea.Msg {
		m.fetcher.Start(context.Background())
		return nil
	})

	// refresh at start
	cmds = append(cmds, func() tea.Msg {
		return m.Refresh(false)
	})

	return tea.Batch(cmds...)
}
