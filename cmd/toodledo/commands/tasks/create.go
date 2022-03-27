package tasks

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg"
	"github.com/alswl/go-toodledo/pkg/models"
	tpriority "github.com/alswl/go-toodledo/pkg/models/enums/tasks/priority"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cmdCreateQuery present the parameters for the command
// parse query with cmd
type cmdCreateQuery struct {
	Context  string
	Folder   string
	Goal     string
	Priority string `validate:"omitempty,oneof=Top top High high Medium medium Low low Negative negative"`
	Status   string `validate:"omitempty,oneof=None NextAction Active Planning Delegated Waiting Hold Postponed Someday Canceled Reference none nextaction active planning delegated waiting hold postponed someday canceled reference"`

	DueDate string `validate:"omitempty,datetime=2006-01-02" json:"due_date" description:"format 2021-01-01"`
	// TODO
	// Tags
}

func (q *cmdCreateQuery) ToQuery(contextSvc services.ContextService, folderSvc services.FolderService,
	goalSvc services.GoalService) (*queries.TaskCreateQuery, error) {
	var err error
	var query queries.TaskCreateQuery

	if q.Context != "" {
		context, err := contextSvc.Find(q.Context)
		if err != nil {
			return nil, errors.Wrap(err, "failed to find context")
		}
		query.ContextID = context.ID
	}
	if q.Folder != "" {
		folder, err := folderSvc.Find(q.Folder)
		if err != nil {
			return nil, errors.Wrap(err, "failed to find folder")
		}
		query.FolderID = folder.ID
	}
	if q.Goal != "" {
		goal, err := goalSvc.Find(q.Goal)
		if err != nil {
			return nil, errors.Wrap(err, "failed to find goal")
		}
		query.GoalID = goal.ID
	}
	query.DueDate = q.DueDate
	query.Priority = tpriority.PriorityString2Type(q.Priority)

	return &query, err
}

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a task",
	Example: "toodledo tasks create --context=1 --folder=2 --goal=3 --priority=High --due_date=2020-01-01 title",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmdQ := cmdCreateQuery{}
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
		taskRichSvc, err := injector.InitTaskRichService()
		if err != nil {
			logrus.WithError(err).Fatal("init task rich service failed")
			return
		}
		contextSvc, err := injector.InitContextCachedService()
		if err != nil {
			logrus.Fatal(err)
			return
		}
		folderSvc, err := injector.InitFolderCachedService()
		if err != nil {
			logrus.Fatal(err)
			return
		}
		goalSvc, err := injector.InitGoalCachedService()
		if err != nil {
			logrus.Fatal(err)
			return
		}
		q, err := cmdQ.ToQuery(contextSvc, folderSvc, goalSvc)
		if err != nil {
			logrus.WithError(err).Fatal("parse query failed")
		}
		q.Title = args[0]

		// TODO simple worked
		t, err := svc.CreateByQuery(q)
		if err != nil {
			logrus.WithError(err).Fatal("create task failed")
			return
		}
		rts, err := taskRichSvc.RichThem([]*models.Task{t})
		fmt.Println(render.Tables4RichTasks(rts))
	},
}

func init() {
	err := pkg.BindFlagsByQuery(createCmd, cmdCreateQuery{})
	if err != nil {
		panic(errors.Wrapf(err, "failed to generate flags for command %s", createCmd.Use))
	}
	//createCmd.Flags().String("title", "", "title of the task")
}
