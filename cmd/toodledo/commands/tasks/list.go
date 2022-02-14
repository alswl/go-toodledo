package tasks

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg"
	tpriority "github.com/alswl/go-toodledo/pkg/models/enums/tasks/priority"
	tstatus "github.com/alswl/go-toodledo/pkg/models/enums/tasks/status"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type cmdListQuery struct {
	ContextID int64
	Context   string
	// FIXME name
	FolderID int64
	// FIXME name
	GoalID   int64
	Priority string `validate:"omitempty,oneof=Top top High high Medium medium Low low Negative negative"`
	Status   string `validate:"omitempty,oneof=None NextAction Active Planning Delegated Waiting Hold Postponed Someday Canceled Reference none nextaction active planning delegated waiting hold postponed someday canceled reference"`

	DueDate string `validate:"omitempty,datetime=2006-01-02" json:"due_date" description:"format 2021-01-01"`
	// TODO
	// Tags
}

func (q *cmdListQuery) PrepareIDs() error {
	if q.Context != "" {
		svc, err := injector.InitContextCachedService()
		if err != nil {
			return errors.Wrap(err, "failed to init folder service")
		}
		// TODO case sensitive
		c, err := svc.Find(q.Context)
		if err != nil {
			return errors.Wrap(err, "failed to get context by name")
		}
		q.ContextID = c.ID
	}
	return nil
}

func (q *cmdListQuery) ToQuery() (*queries.TaskListQuery, error) {
	var query = &queries.TaskListQuery{}

	query.ContextID = q.ContextID
	query.FolderID = q.FolderID
	query.GoalID = q.GoalID
	query.DueDate = q.DueDate
	if q.Priority != "" {
		p := tpriority.PriorityString2Type(q.Priority)
		query.Priority = &p
	}
	if q.Status != "" {
		s := tstatus.StatusString2Type(q.Status)
		query.Status = &s
	}

	return query, nil
}

var listCmd = &cobra.Command{
	Use:  "list",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmdQ := cmdListQuery{}
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
		svc, err := injector.InitTaskCachedService()
		if err != nil {
			logrus.WithError(err).Fatal("failed to init task service")
			return
		}
		syncer, err := injector.InitSyncer()
		if err != nil {
			logrus.WithError(err).Fatal("init syncer failed")
			return
		}
		taskRichSvc, err := injector.InitTaskRichService()
		if err != nil {
			logrus.WithError(err).Fatal("init task rich service failed")
			return
		}
		err = syncer.SyncOnce()
		if err != nil {
			logrus.WithError(err).Fatal("sync failed")
			return
		}
		err = cmdQ.PrepareIDs()
		if err != nil {
			logrus.WithError(err).Fatal("prepare ids failed")
			return
		}
		q, err := cmdQ.ToQuery()
		if err != nil {
			logrus.WithError(err).Fatal("parse query failed")
		}

		tasks, err := svc.ListAllByQuery(q)
		if err != nil {
			logrus.Error(err)
			return
		}
		rts, err := taskRichSvc.RichThem(tasks)
		fmt.Println(render.Tables4RichTasks(rts))
	},
}

func init() {
	err := pkg.BindFlagsByQuery(listCmd, cmdListQuery{})
	if err != nil {
		panic(errors.Wrapf(err, "failed to generate flags for command %s", listCmd.Use))
	}
}
