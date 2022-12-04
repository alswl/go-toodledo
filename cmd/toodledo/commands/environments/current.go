package environments

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCurrentCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "current",
		Args:  cobra.NoArgs,
		Short: "Show current environment",
		Example: heredoc.Doc(`
			$ toodledo environment current
`),
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
			} else {
				_, _ = fmt.Fprintf(f.IOStreams.Out, "%s, %s\n", matched.Name, matched.Project)
			}
		},
	}
}
