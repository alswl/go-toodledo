package tasks

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/models"
	tpriority "github.com/alswl/go-toodledo/pkg/models/enums/tasks/priority"
	tstatus "github.com/alswl/go-toodledo/pkg/models/enums/tasks/status"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/alswl/go-toodledo/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type cmdCreateQuery struct {
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

func (q *cmdCreateQuery) ToQuery(contextSvc services.ContextService, folderSvc services.FolderService,
	goalSvc services.GoalService) (*queries.TaskCreateQuery, error) {
	query := queries.TaskCreateQuery{}

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
	query.Priority = tpriority.PriorityString2Type(q.Priority)
	query.Status = tstatus.StatusString2Type(q.Status)
	query.DueDate = q.DueDate
	query.Title = q.Title

	return &query, nil
}

var createCmd = &cobra.Command{
	Use:     "create",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"new"},
	Short:   "Create a task",
	Example: heredoc.Doc(`
		$ toodledo tasks create --context=1 --folder=2 --goal=3 --priority=High --due_date=2020-01-01 title
`),
	Run: func(cmd *cobra.Command, args []string) {
		cmdQ := cmdCreateQuery{}
		err := utils.FillQueryByFlags(cmd, &cmdQ)
		if err != nil {
			logrus.WithError(err).Fatal("failed")
		}
		validate := validator.New()
		err = validate.Struct(cmdQ)
		if err != nil {
			logrus.WithError(err).Fatal("validate failed")
		}

		app, err := injector.InitCLIApp()
		if err != nil {
			logrus.Fatal("login required, using `toodledo auth login` to login.")
			return
		}
		svc := app.TaskSvc
		taskRichSvc := app.TaskRichSvc
		contextSvc := app.ContextSvc
		folderSvc := app.FolderSvc
		goalSvc := app.GoalSvc
		cmdQ.Title = args[0]
		q, err := cmdQ.ToQuery(contextSvc, folderSvc, goalSvc)
		if err != nil {
			logrus.WithError(err).Fatal("parse query failed")
		}

		// TODO simple worked
		t, err := svc.CreateByQuery(q)
		if err != nil {
			logrus.WithError(err).Fatal("create task failed")
			return
		}
		rt, _ := taskRichSvc.Rich(t)
		fmt.Println(render.Tables4RichTasks([]*models.RichTask{rt}))
	},
}

func NewCreateCmd(f *cmdutil.Factory) *cobra.Command {
	cmd := createCmd
	err := utils.BindFlagsByQuery(cmd, cmdCreateQuery{})
	if err != nil {
		logrus.WithError(err).Fatal("bind flags failed")
	}
	return cmd
}

func init() {
}
