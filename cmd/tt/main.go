package main

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/tt/app"
	"github.com/alswl/go-toodledo/pkg/common/logging"
	"github.com/alswl/go-toodledo/pkg/models"
	tstatus "github.com/alswl/go-toodledo/pkg/models/enums/tasks/status"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
)

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
	viper.AddConfigPath(path.Join(home, ".config", "toodledo"))
	viper.SetConfigType("yaml")
	viper.SetConfigName(".toodledo")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		logrus.Debug("config file", viper.ConfigFileUsed())
	}
}

func main() {
	// TODO
	initViper()
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
