package tasks

import (
	"fmt"
	"strconv"

	"github.com/alswl/go-toodledo/pkg/utils/flags"

	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/models"
	tpriority "github.com/alswl/go-toodledo/pkg/models/enums/tasks/priority"
	tstatus "github.com/alswl/go-toodledo/pkg/models/enums/tasks/status"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/thoas/go-funk"
)

// cmdEditQuery contains the desired parameters for the task.
// nolint: lll // ignore long line length
type cmdEditQuery struct {
	// TODO
	// ContextID int64
	Context string
	// TODO
	// FolderID int64
	Folder string
	// TODO
	// GoalID int64
	Goal     string
	Priority string `validate:"omitempty,oneof=Top top High high Medium medium Low low Negative negative"`
	Status   string `validate:"omitempty,oneof=None NextAction Active Planning Delegated Waiting Hold Postponed Someday Canceled Reference none nextaction active planning delegated waiting hold postponed someday canceled reference"`

	DueDate string `validate:"omitempty,datetime=2006-01-02" json:"due_date" description:"format 2021-01-01"`
	// TODO
	// Tags
	Title string
}

func toQuery(contextSvc services.ContextService, folderSvc services.FolderService,
	goalSvc services.GoalService, q *cmdEditQuery) (*queries.TaskEditQuery, error) {
	query := queries.TaskEditQuery{}

	if q.Context != "" {
		context, err := contextSvc.Find(q.Context)
		if err != nil {
			return nil, fmt.Errorf("find context failed: %w", err)
		}
		query.ContextID = context.ID
	}
	if q.Folder != "" {
		folder, err := folderSvc.Find(q.Folder)
		if err != nil {
			return nil, fmt.Errorf("find folder failed: %w", err)
		}
		query.FolderID = folder.ID
	}
	if q.Goal != "" {
		goal, err := goalSvc.Find(q.Goal)
		if err != nil {
			return nil, fmt.Errorf("find goal failed: %w", err)
		}
		query.GoalID = goal.ID
	}
	query.Priority = tpriority.String2Type(q.Priority)
	query.Status = tstatus.String2Type(q.Status)
	query.DueDate = q.DueDate
	if q.Title != "" {
		query.Title = q.Title
	}
	return &query, nil
}

func NewEditCmd(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit",
		Args:  cobra.ExactArgs(1),
		Short: "Edit a task",
		Example: heredoc.Doc(`
			$ toodledo tasks edit 8848
			$ toodledo tasks edit --title="New title" 8848
			$ toodledo tasks edit --context=Work 8848
			$ toodledo tasks edit --folder=Inbox 8848
			$ toodledo tasks edit --goal=landing-moon 8848
			$ toodledo tasks edit --priority=High 8848
			$ toodledo tasks edit --status=Active 8848
		`),
		Run: func(cmd *cobra.Command, args []string) {
			cmdQ := cmdEditQuery{}
			err := flags.FillQueryByFlags(cmd, &cmdQ)
			if err != nil {
				logrus.WithError(err).Fatal("parse query failed")
			}
			if funk.IsZero(cmdQ) {
				logrus.Fatal("query is empty")
				return
			}
			// services
			app, err := injector.InitCLIApp()
			if err != nil {
				logrus.WithError(err).Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			taskSvc := app.TaskSvc
			contextSvc := app.ContextSvc
			folderSvc := app.FolderSvc
			goalSvc := app.GoalSvc
			taskRichSvc := app.TaskRichSvc

			// fetch task
			id, _ := strconv.Atoi(args[0])
			_, err = taskSvc.FindByID(int64(id))
			if err != nil {
				logrus.WithError(err).Fatal("find task")
				return
			}

			// query
			q, err := toQuery(contextSvc, folderSvc, goalSvc, &cmdQ)
			if err != nil {
				logrus.WithError(err).Fatal("parse query failed")
				return
			}
			q.ID = int64(id)

			newT, err := taskSvc.EditByQuery(q)
			if err != nil {
				logrus.WithError(err).Fatal("edit task")
			}

			rt, _ := taskRichSvc.Rich(newT)
			_, _ = fmt.Fprintln(f.IOStreams.Out, render.Tables4RichTasks([]*models.RichTask{rt}))
		},
	}
	err := flags.BindFlagsByQuery(cmd, cmdEditQuery{})
	if err != nil {
		logrus.WithError(err).Fatal("bind flags failed")
		return nil
	}
	return cmd
}
