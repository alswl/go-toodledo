package folders

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/folder"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:  "create",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		auth, err := client.ProvideSimpleAuth()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		name := args[0]

		cli := client.NewHTTPClient(strfmt.NewFormats())
		params := folder.NewPostFoldersAddPhpParams()
		params.SetName(name)
		res, err := cli.Folder.PostFoldersAddPhp(params, auth)
		if err != nil {
			logrus.Error(err)
			return
		}
		fmt.Print(render.Tables4Folder(res.Payload))
	},
}
