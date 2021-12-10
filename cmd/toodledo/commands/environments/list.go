package environments

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		cSrv := services.NewEnvironmentService()
		cks, err := cSrv.ListAll()
		if err != nil {
			logrus.Error(err)
			return
		}

		fmt.Printf(render.RenderEnvironments(cks))
	},
}
