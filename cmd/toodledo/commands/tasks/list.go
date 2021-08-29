package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/alswl/go-toodledo/pkg/auth"
	"github.com/alswl/go-toodledo/pkg/client"
	"github.com/alswl/go-toodledo/pkg/client/task"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/go-openapi/strfmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		auth, err := auth.ProvideSimpleAuth()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		cli := client.NewHTTPClient(strfmt.NewFormats())

		params := task.NewGetTasksGetPhpParams()
		comp := int64(0)
		params.SetComp(&comp)
		num := int64(10)
		params.SetNum(&num)
		res, err := cli.Task.GetTasksGetPhp(params, auth)
		if err != nil {
			logrus.Error(err)
			return
		}
		var info models.PaginatedInfo
		var tasks []*models.Task
		for i, x := range res.Payload {
			if i == 0 {
				bytes, _ := json.Marshal(x)
				// TODO using service
				json.Unmarshal(bytes, &info)
				continue
			}
			bytes, _ := json.Marshal(x)
			var t models.Task
			// TODO using service
			json.Unmarshal(bytes, &t)
			tasks = append(tasks, &t)
		}
		fmt.Print(info)
		fmt.Print(render.Tables4Task(tasks))
	},
}
