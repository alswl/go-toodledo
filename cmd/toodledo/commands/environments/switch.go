package environments

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewSwitchCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "switch",
		Args:  cobra.ExactArgs(1),
		Short: "Switch environment",
		Example: heredoc.Doc(`
			$ toodledo environment switch <name>
		`),
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if len(args) != 0 {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			cSrv := services.NewEnvironmentService()
			keys, err := cSrv.ListAllKeys()
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
			_, _ = fmt.Fprintln(f.IOStreams.Out, "Switch successfully")
		},
	}
}
