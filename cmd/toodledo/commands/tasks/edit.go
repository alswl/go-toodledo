package tasks

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/models"
	tpriority "github.com/alswl/go-toodledo/pkg/models/enums/tasks/priority"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/alswl/go-toodledo/pkg/utils"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"
	"strconv"
)

// cmdEditQuery contains the desired parameters for the task
type cmdEditQuery struct {
	// TODO
	//ContextID int64
	Context string
	// TODO
	//FolderID int64
	Folder string
	// TODO
	//GoalID int64
	Goal     string
	Priority string `validate:"omitempty,oneof=Top top High high Medium medium Low low Negative negative"`
	Status   string `validate:"omitempty,oneof=None NextAction Active Planning Delegated Waiting Hold Postponed Someday Canceled Reference none nextaction active planning delegated waiting hold postponed someday canceled reference"`

	DueDate string `validate:"omitempty,datetime=2006-01-02" json:"due_date" description:"format 2021-01-01"`
	// TODO
	// Tags
	Title string
}

func (q *cmdEditQuery) ToQuery(contextSvc services.ContextService, folderSvc services.FolderService,
	goalSvc services.GoalService) (*queries.TaskEditQuery, error) {
	query := queries.TaskEditQuery{}

	if q.Context != "" {
		context, err := contextSvc.Find(q.Context)
		if err != nil {
			return nil, errors.Wrap(err, "find context")
		}
		query.ContextID = context.ID
	}
	if q.Folder != "" {
		folder, err := folderSvc.Find(q.Folder)
		if err != nil {
			return nil, errors.Wrap(err, "find folder")
		}
		query.FolderID = folder.ID
	}
	if q.Goal != "" {
		goal, err := goalSvc.Find(q.Goal)
		if err != nil {
			return nil, errors.Wrap(err, "find goal")
		}
		query.GoalID = goal.ID
	}
	query.DueDate = q.DueDate
	query.Priority = tpriority.PriorityString2Type(q.Priority)
	if q.Title != "" {
		query.Title = q.Title
	}
	return &query, nil
}

var editCmd = &cobra.Command{
	Use:  "edit",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmdQ := cmdEditQuery{}
		err := utils.FillQueryByFlags(cmd, &cmdQ)
		if err != nil {
			logrus.WithError(err).Fatal("parse query failed")
		}
		// services
		_, err = injector.InitApp()
		if err != nil {
			logrus.WithError(err).Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		svc, err := injector.InitTaskService()
		if err != nil {
			logrus.WithError(err).Fatal("init task service")
			return
		}
		taskRichSvc, err := injector.InitTaskRichService()
		if err != nil {
			logrus.WithError(err).Error("init task rich service")
			return
		}
		contextSvc, err := injector.InitContextCachedService()
		if err != nil {
			logrus.WithError(err).Fatal("init context service")
			return
		}
		folderSvc, err := injector.InitFolderCachedService()
		if err != nil {
			logrus.WithError(err).Fatal("init folder service")
			return
		}
		goalSvc, err := injector.InitGoalCachedService()
		if err != nil {
			logrus.WithError(err).Fatal("init goal service")
			return
		}

		// fetch task
		id, _ := strconv.Atoi(args[0])
		_, err = svc.FindById(int64(id))
		if err != nil {
			logrus.WithError(err).Fatal("find task")
			return
		}

		// query
		q, err := cmdQ.ToQuery(contextSvc, folderSvc, goalSvc)
		if err != nil {
			logrus.WithError(err).Fatal("parse query failed")
		}
		// FIXME not works
		if funk.IsZero(q) {
			logrus.Fatal("query is empty")
			return
		}
		q.ID = int64(id)

		newT, err := svc.EditByQuery(q)
		if err != nil {
			logrus.WithError(err).Fatal("edit task")
		}

		// FIXME rich is cached service, using it with params
		rt, _ := taskRichSvc.Rich(newT)
		fmt.Println(render.Tables4RichTasks([]*models.RichTask{rt}))
	},
}

func init() {
	err := utils.BindFlagsByQuery(editCmd, cmdEditQuery{})
	if err != nil {
		panic(errors.Wrapf(err, "generate flags for command %s", editorCmd.Use))
	}

	TaskCmd.AddCommand(editCmd)
}
