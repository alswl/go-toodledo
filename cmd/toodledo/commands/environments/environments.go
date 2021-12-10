package environments

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/constants"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"
)

const DefaultEnvironmentKey = "default-environment"

var EnvironmentsCmd = &cobra.Command{
	Use:   "environment",
	Short: "Switch for the variable reality",
}

func init() {
	switchCmd.RegisterFlagCompletionFunc(constants.ArgEnvironment, func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		cSrv := services.NewEnvironmentService()
		cks, err := cSrv.ListAll()
		if err != nil {
			logrus.Warn(err)
			return []string{}, cobra.ShellCompDirectiveDefault
		}
		keys := funk.Map(cks, func(x *models.EnvironmentWithKey) string {
			// TODO using description, v2 completions
			//return fmt.Sprintf("%s", x.Key, x.Name)
			return fmt.Sprintf("%s", x.Key)
		}).([]string)

		return keys, cobra.ShellCompDirectiveNoFileComp
	})

	EnvironmentsCmd.AddCommand(listCmd, switchCmd, currentCmd)
}
