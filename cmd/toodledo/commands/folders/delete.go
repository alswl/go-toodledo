package folders

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/folder"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:  "delete",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		auth, err := client.ProvideSimpleAuth()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		name := args[0]

		f, err := services.FindFolderByName(auth, name)
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
