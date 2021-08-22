package folders

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/auth"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/alswl/go-toodledo/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var UnArchiveCmd = &cobra.Command{
	Use:  "unarchive",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		t := viper.GetString("auth.access_token")
		if t == "" {
			logrus.Error("auth.access_token is empty")
			return
		}
		auth := auth.NewSimpleAuth(t)
		name := args[0]

		f, err := service.FindFolderByName(auth, name)
		if err != nil {
			logrus.Error(err)
			return
		}

		newF, err := service.ArchiveFolder(auth, int(f.ID), false)
		if err != nil {
			logrus.Error(err)
			return
		}

		fmt.Print(render.TablesRender([]*models.Folder{newF}))
	},
}
