package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type FilterFormModel struct {
	context []string
	status  []string
	goal    []string
}

func (m FilterFormModel) Init() tea.Cmd {
	return nil
}

func (m FilterFormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m FilterFormModel) View() string {
	cs := strings.Join(m.context, ",")
	ss := strings.Join(m.status, ",")
	gs := strings.Join(m.goal, ",")

	return fmt.Sprintf("Context: %s Status: %s Goal: %s", cs, ss, gs)
}
