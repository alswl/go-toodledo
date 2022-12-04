package main

import (
	"os"

	"github.com/alswl/go-toodledo/cmd/tt/app"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/common/logging"
	tea "github.com/charmbracelet/bubbletea"
)

func init() {
	common.InitViper("", ".config/toodledo", "conf")
}

func main() {
	err := logging.InitFactory("/tmp/tt", false, false)
	log := logging.ProvideLogger()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// TODO full screen
	p := tea.NewProgram(app.InitialModel(), tea.WithAltScreen())
	if err = p.Start(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
