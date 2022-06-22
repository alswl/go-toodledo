package main

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/tt/app"
	"github.com/alswl/go-toodledo/pkg/common/logging"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
)

func initViper() {
	// Find home directory.
	home, err := os.UserHomeDir()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to get user home directory")
		return
	}

	// Search config in home directory with name ".toodledo" (without extension).
	viper.AddConfigPath(path.Join(home, ".config", "toodledo"))
	viper.SetConfigType("yaml")
	viper.SetConfigName("toodledo")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		logrus.Debug("config file", viper.ConfigFileUsed())
	}
}

func main() {
	err := logging.InitFactory("/tmp/tt", false, false)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	initViper()

	// TODO full screen
	p := tea.NewProgram(app.InitialModel(), tea.WithAltScreen())
	if err = p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
