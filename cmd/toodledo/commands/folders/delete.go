package folders

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/auth"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/folder"
	"github.com/alswl/go-toodledo/pkg/service"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var DeleteCmd = &cobra.Command{
	Use:  "delete",
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

		cli := client.NewHTTPClient(strfmt.NewFormats())
		params := folder.NewPostFoldersDeletePhpParams()
		params.SetID(f.ID)
		resp, err := cli.Folder.PostFoldersDeletePhp(params, auth)
		if err != nil {
			logrus.WithField("resp", resp).Error(err)
			return
		}
		fmt.Print("done")
	},
}
