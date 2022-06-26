package config

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func NewViewCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "view",
		Args:  cobra.ExactArgs(0),
		Short: "View config",
		Example: heredoc.Doc(`
			$ toodledo config view
`),
		Run: func(cmd *cobra.Command, args []string) {
			settings := viper.AllSettings()
			bs, _ := yaml.Marshal(settings)
			fmt.Println(string(bs))
		},
	}
}
