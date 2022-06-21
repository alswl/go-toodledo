package config

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func NewViewCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use: "view",
		Run: func(cmd *cobra.Command, args []string) {
			settings := viper.AllSettings()
			text, _ := yaml.Marshal(settings)
			fmt.Println(text)
		},
	}
}
