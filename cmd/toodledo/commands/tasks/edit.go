package tasks

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

var EditCmd = &cobra.Command{
	Use:  "edit",
	Args: cobra.ExactArgs(1),
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

		id, _ := strconv.Atoi(args[0])
		newT, err := svc.Edit(id, &models.Task{})
		if err != nil {
			logrus.Fatal(err)
			return
		}
		// FIXME
		fmt.Println(newT)
	},
}
