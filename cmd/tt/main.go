package main

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/tt/app"
	"github.com/alswl/go-toodledo/pkg/common"
	"github.com/alswl/go-toodledo/pkg/common/logging"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func init() {
	common.InitViper("", ".config/toodledo", "conf")
}

func main() {
	err := logging.InitFactory("/tmp/tt", false, false)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	// TODO full screen
	p := tea.NewProgram(app.InitialModel(), tea.WithAltScreen())
	if err = p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
