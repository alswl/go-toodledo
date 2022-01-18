package tasks

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg"
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/enums/tasks"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

// cmdQuery present the parameters for the command
// parse query with cmd
type cmdQuery struct {
	ContextID int64
	FolderID  int64
	GoalID    int64
	// XXX
	Priority string `validate:"oneof=Top High Medium Low Negative"`

	DueDate time.Time `json:"due_date"`
	// XXX
	// Tags
}

func (q *cmdQuery) ToQuery() (*queries.TaskCreateQuery, error) {
	var err error
	var query queries.TaskCreateQuery

	query.ContextID = q.ContextID
	query.FolderID = q.FolderID
	query.GoalID = q.GoalID
	query.DueDate = q.DueDate
	query.Priority = tasks.PriorityString2Type(q.Priority)

	return &query, err
}

var CreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a task",
	Example: "toodledo tasks create --context=1 --folder=2 --goal=3 --priority=High --due_date=2020-01-01 title",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmdQ := cmdQuery{}
		err := pkg.FillQueryByFlags(cmd, &cmdQ)
		if err != nil {
			logrus.WithError(err).Fatal("failed")
		}
		validate := validator.New()
		err = validate.Struct(cmdQ)
		if err != nil {
			logrus.WithError(err).Fatal("validate failed")
		}

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
		q, err := cmdQ.ToQuery()
		if err != nil {
			logrus.WithError(err).Fatal("parse query failed")
		}
		q.Title = args[0]

		// TODO simple worked
		t, err := svc.CreateWithQuery(q)
		if err != nil {
			logrus.WithError(err).Fatal("create task failed")
			return
		}

		fmt.Println(render.Tables4Task([]*models.Task{t}))
	},
}

func init() {
	err := pkg.BindFlagsByQuery(CreateCmd, cmdQuery{})
	if err != nil {
		panic(errors.Wrapf(err, "failed to generate flags for command %s", CreateCmd.Use))
	}
	//CreateCmd.Flags().String("title", "", "title of the task")
}
