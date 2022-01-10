package tasks

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg"
	"github.com/alswl/go-toodledo/pkg/models/enums"
	"github.com/alswl/go-toodledo/pkg/models/enums/tasks"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
	"time"
)

// Query holds the query parameters for the command
// parse query with cmd
//type Query struct {
//	Title     string `json:"title,omitempty" validate:"required"`
//	ContextID int64  `json:"context_id,omitempty"`
//	FolderID  int64  `json:"folder_id,omitempty"`
//	GoalID    int64  `json:"goal_id,omitempty"`
//
//	DueDate time.Time `json:"due_date"`
//}

var CreateCmd = &cobra.Command{
	Use:  "create",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		q := queries.TaskCreateQuery{}
		err := pkg.FillQueryByStructuredCmd(cmd, &q)
		// XXX
		//validate q
		if err != nil {
			logrus.WithError(err).Fatal("failed")
		}

		qb := queries.NewTaskCreateQueryBuilder()

		// filled qb with corba arguments
		s, err := cmd.Flags().GetString(string(enums.TASK_FIELD_TITLE))
		qb.WithTitle(s)
		// TODO
		//s, err = cmd.Flags().GetString(string(enums.TASK_FIELD_FOLDER))
		//qb.WithFolder(s)
		//s, err = cmd.Flags().GetString(string(enums.TASK_FIELD_CONTEXT))
		//qb.WithContext(s)
		b, err := cmd.Flags().GetBool(string(enums.TASK_FIELD_STAR))
		qb.WithStar(b)
		i, err := cmd.Flags().GetInt64(string(enums.TASK_FIELD_PRIORITY))
		qb.WithPriority(tasks.PriorityValue2Type(i))
		s, err = cmd.Flags().GetString(string(enums.TASK_FIELD_DUE_DATE))
		t, err := time.Parse("2006-01-02", s)
		qb.WithDueDate(t)
		s, err = cmd.Flags().GetString(string(enums.TASK_FIELD_DUE_TIME))
		splits := strings.SplitN(s, ":", 2)
		h, _ := strconv.Atoi(splits[0])
		m, _ := strconv.Atoi(splits[1])
		d := time.Date(1970, 1, 1, h, m, 0, 0, time.UTC)
		qb.WithDueTime(d.Unix())

		_, err = injector.InitApp()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}

		title := args[0]

		_, err = injector.InitTaskService()
		if err != nil {
			logrus.Fatal(err)
			return
		}
		_ = qb.WithTitle(title).Build()
		//validate

		//t, err := svc.CreateWithQuery()

		if err != nil {
			logrus.Fatal(err)
			return
		}

		//fmt.Println(render.Tables4Task([]*models.Task{t}))
	},
}

func init() {
	err := pkg.GenerateFlagsByStructure(CreateCmd, queries.TaskCreateQuery{})
	if err != nil {
		panic(err)
	}
	//CreateCmd.Flags().String("title", "", "title of the task")
}
