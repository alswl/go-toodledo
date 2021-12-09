package environments

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var currentCmd = &cobra.Command{
	Use: "current",
	Run: func(cmd *cobra.Command, args []string) {
		name := viper.GetString(DefaultEnvironmentKey)
		if name == "" {
			logrus.Error("not set")
			return
		}
		var cs map[string]models.Environment
		err := viper.UnmarshalKey("environments", &cs)
		if err != nil {
			logrus.Error(err)
			return
		}
		if matched, ok := cs[name]; !ok {
			logrus.WithField("name", name).Error("not found")
			return
		} else {
			fmt.Printf("%s, %s\n", matched.Name, matched.Project)
		}
	},
}
