package commands

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("%s.%s.%s-%s %s", pkg.MajorVersion, pkg.MinorVersion,
			pkg.PatchVersion, pkg.BuildVersion, pkg.Package))
	},
}
