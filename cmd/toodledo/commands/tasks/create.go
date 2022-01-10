package tasks

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

// cmdQuery holds the query parameters for the command
// parse query with cmd
type cmdQuery struct {
	Title     string `json:"title,omitempty" validate:"required"`
	ContextID int64  `json:"context_id,omitempty"`
	FolderID  int64  `json:"folder_id,omitempty"`
	GoalID    int64  `json:"goal_id,omitempty"`

	DueDate time.Time `json:"due_date"`
}

func (q *cmdQuery) ToQuery() (*queries.TaskCreateQuery, error) {
	var err error
	var query queries.TaskCreateQuery

	query.Title = q.Title
	query.ContextID = q.ContextID
	query.FolderID = q.FolderID
	query.GoalID = q.GoalID
	query.DueDate = q.DueDate

	return &query, err
}

var CreateCmd = &cobra.Command{
	Use:  "create",
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		//q := queries.TaskCreateQuery{}
		q := cmdQuery{}
		err := pkg.FillQueryByFlags(cmd, &q)
		//validate q
		if err != nil {
			logrus.WithError(err).Fatal("failed")
		}

		//qb := queries.NewTaskCreateQueryBuilder()
		// filled qb with corba arguments
		//s, err := cmd.Flags().GetString(string(enums.TASK_FIELD_TITLE))
		//qb.WithTitle(s)
		// TODO
		//s, err = cmd.Flags().GetString(string(enums.TASK_FIELD_FOLDER))
		//qb.WithFolder(s)
		//s, err = cmd.Flags().GetString(string(enums.TASK_FIELD_CONTEXT))
		//qb.WithContext(s)
		//b, err := cmd.Flags().GetBool(string(enums.TASK_FIELD_STAR))
		//qb.WithStar(b)
		//i, err := cmd.Flags().GetInt64(string(enums.TASK_FIELD_PRIORITY))
		//qb.WithPriority(tasks.PriorityValue2Type(i))
		//s, err = cmd.Flags().GetString(string(enums.TASK_FIELD_DUE_DATE))
		//t, err := time.Parse("2006-01-02", s)
		//qb.WithDueDate(t)
		// TODO layout to day and hour
		//s, err = cmd.Flags().GetString(string(enums.TASK_FIELD_DUE_TIME))
		//splits := strings.SplitN(s, ":", 2)
		//h, _ := strconv.Atoi(splits[0])
		//m, _ := strconv.Atoi(splits[1])
		//d := time.Date(1970, 1, 1, h, m, 0, 0, time.UTC)
		//qb.WithDueTime(d.Unix())

		_, err = injector.InitApp()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}

		svc, err := injector.InitTaskService()
		if err != nil {
			logrus.Fatal(err)
			return
		}
		//q2 := qb.WithTitle(q.Title).
		//	WithContextID(q.ContextID).
		//	WithFolderID(q.FolderID).
		//	WithGoalID(q.GoalID).
		//	WithDueDate(q.DueDate).
		//	//WithPriority(tasks.PriorityValue2Type(i)).
		//	//WithStar(b)
		//	Build()
		//validate
		q2, err := q.ToQuery()
		if err != nil {
			logrus.WithError(err).Fatal("parse query failed")
		}

		t, err := svc.CreateWithQuery(q2)
		if err != nil {
			logrus.Fatal(err)
			return
		}

		fmt.Println(render.Tables4Task([]*models.Task{t}))
	},
}

func init() {
	err := pkg.GenerateFlagsByQuery(CreateCmd, queries.TaskCreateQuery{})
	if err != nil {
		panic(errors.Wrapf(err, "failed to generate flags for command %s", CreateCmd.Use))
	}
	//CreateCmd.Flags().String("title", "", "title of the task")
}
