package environments

import (
	"fmt"

	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/common/terminal"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/constants"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

const DefaultEnvironmentKey = "default-environment"

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "environment",
		Args:  cobra.NoArgs,
		Short: "Manage the environments",
	}
	_ = (&cobra.Command{
		Use:  "switch key",
		Args: cobra.ExactArgs(1),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) != 0 {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			keys, err := terminal.ListAllKeys()
			if err != nil {
				logrus.Warn(err)
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			return keys, cobra.ShellCompDirectiveNoFileComp
		},
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			var cs map[string]models.Environment
			err := viper.UnmarshalKey("environments", &cs)
			if err != nil {
				logrus.Error(err)
				return
			}
			if _, ok := cs[name]; !ok {
				logrus.WithField("name", name).Error("not found")
				return
			}
			viper.Set(DefaultEnvironmentKey, name)
			_ = viper.WriteConfig()
			_, _ = fmt.Fprintln(f.IOStreams.Out, "Done")
		},
	}).RegisterFlagCompletionFunc(constants.ArgEnvironment,
		func(_ *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
			cks, err := terminal.ListAll()
			if err != nil {
				logrus.Warn(err)
				return []string{}, cobra.ShellCompDirectiveDefault
			}
			keys, _ := funk.Map(cks, func(x *models.EnvironmentWithKey) string {
				// TODO using description, v2 completions
				// return fmt.Sprintf("%s", x.Key, x.Name)
				return x.Key
			}).([]string)

			return keys, cobra.ShellCompDirectiveNoFileComp
		})

	cmd.AddCommand(NewListCmd(f), NewSwitchCmd(f), NewCurrentCmd(f))
	return cmd
}
