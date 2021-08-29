package tasks

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/registries"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := registries.InitAuth()
		svc, _ := registries.InitTaskService()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}

		tasks, paging, err := svc.QueryAll()
		if err != nil {
			logrus.Error(err)
			return
		}

		fmt.Println(paging)
		fmt.Println(render.Tables4Task(tasks))
	},
}
