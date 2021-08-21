package folders

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/auth"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/folder"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
	"strconv"
)

var EditCmd = &cobra.Command{
	Use: "edit",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		t := viper.GetString("auth.access_token")
		if t == "" {
			logrus.Error("auth.access_token is empty")
			return
		}
		auth := auth.NewSimpleAuth(t)
		
		name := args[0]
		newName := args[1]

		cli := client.NewHTTPClient(strfmt.NewFormats())
		ts, err := cli.Folder.GetFoldersGetPhp(folder.NewGetFoldersGetPhpParams(), auth)
		if err != nil {
			logrus.Error(err)
			return
		}
		filter := funk.Filter(ts.Payload, func(x *models.Folder) bool {return x.Name == name }) .([]*models.Folder)
		if len(filter) == 0 {
			logrus.Error("not found")
			return
		}
		task := filter[0]

		p := folder.NewPostFoldersEditPhpParams()
		p.SetID(strconv.Itoa(int(task.ID)))
		p.SetName(&newName)
		res, err := cli.Folder.PostFoldersEditPhp(p, auth)
		if err != nil {
			logrus.Error(err)
			return
		}
		fmt.Print(render.TablesRender(res.Payload))
	},
}
