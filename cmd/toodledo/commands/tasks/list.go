package tasks

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:  "list",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := injector.InitApp()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		svc, err := injector.InitTaskService()
		if err != nil {
			logrus.Fatal(err)
			return
		}

		tasks, paging, err := svc.ListAll()
		if err != nil {
			logrus.Error(err)
			return
		}

		fmt.Println(paging)
		fmt.Println(render.Tables4Task(tasks))
	},
}

func init() {
	//err := pkg.BindFlagsByQuery(listCmd, cmdQuery{})
	//if err != nil {
	//	panic(errors.Wrapf(err, "failed to generate flags for command %s", listCmd.Use))
	//}
}
