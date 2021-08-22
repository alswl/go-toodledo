package goals

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/auth"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/goal"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/alswl/go-toodledo/pkg/service"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
)

var RenameCmd = &cobra.Command{
	Use: "rename",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		t := viper.GetString("auth.access_token")
		if t == "" {
			logrus.Error("auth.access_token is empty")
			return
		}
		auth := auth.NewSimpleAuth(t)
		cli := client.NewHTTPClient(strfmt.NewFormats())
		
		name := args[0]
		newName := args[1]
		if name == newName {
			logrus.Error("not changed")
			return
		}

		f, err := service.FindGoalByName(auth, name)
		if err != nil {
			logrus.Error(err)
			return
		}

		p := goal.NewPostGoalsEditPhpParams()
		p.SetID(strconv.Itoa(int(f.ID)))
		p.SetName(&newName)
		res, err := cli.Goal.PostGoalsEditPhp(p, auth)
		if err != nil {
			logrus.Error(err)
			return
		}
		fmt.Print(render.Tables4Goal(res.Payload))
	},
}
