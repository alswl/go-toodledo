package folders

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/folder"
	"github.com/alswl/go-toodledo/pkg/registries"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		auth, err := registries.InitAuth()
		if err != nil {
			logrus.Error(err)
			return
		}

		cli := client.NewHTTPClient(strfmt.NewFormats())
		res, err := cli.Folder.GetFoldersGetPhp(folder.NewGetFoldersGetPhpParams(), auth)
		if err != nil {
			logrus.Error(err)
			return
		}
		fmt.Print(render.Tables4Folder(res.Payload))
	},
}
