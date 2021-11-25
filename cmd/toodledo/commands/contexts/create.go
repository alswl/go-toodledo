package contexts

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:  "create",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app, err := injector.InitApp()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		name := args[0]

		created, err := app.ContextSvc.Create(name)
		if err != nil {
			logrus.Error(err)
			return
		}

		fmt.Println(render.Tables4Context([]*models.Context{created}))
	},
}
