package main

import (
	"os"

	"github.com/sirupsen/logrus"

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
	stdLog := logrus.StandardLogger()
	if err != nil {
		stdLog.WithError(err).Error("failed to init logging factory")
		os.Exit(1)
	}

	// TODO full screen
	mainModel, err := app.InitialModel()
	if err != nil {
		stdLog.WithError(err).Fatal("failed to initialize model")
		os.Exit(1)
	}
	p := tea.NewProgram(mainModel, tea.WithAltScreen())
	if _, ierr := p.Run(); ierr != nil {
		stdLog.WithError(ierr).Fatal("failed to start program")
		os.Exit(1)
	}
}
