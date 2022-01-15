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
	Title     string `validate:"required"`
	ContextID int64
	FolderID  int64
	GoalID    int64

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
		cmdQ := cmdQuery{}
		err := pkg.FillQueryByFlags(cmd, &cmdQ)
		if err != nil {
			logrus.WithError(err).Fatal("failed")
		}
		// TODO validate by validator

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
