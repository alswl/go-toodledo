package folders

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:  "delete",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := injector.InitApp()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		svc, err := injector.InitFolderService()
		if err != nil {
			logrus.Fatal(err)
			return
		}

		name := args[0]
		err = svc.Delete(name)
		if err != nil {
			logrus.Fatal(err)
			return
		}
		fmt.Println("done")
	},
}
