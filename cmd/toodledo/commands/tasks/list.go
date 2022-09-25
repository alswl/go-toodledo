package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/fetchers"
	tpriority "github.com/alswl/go-toodledo/pkg/models/enums/tasks/priority"
	tstatus "github.com/alswl/go-toodledo/pkg/models/enums/tasks/status"
	"github.com/alswl/go-toodledo/pkg/models/enums/tasks/subtasksview"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/alswl/go-toodledo/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"sigs.k8s.io/yaml"
)

type cmdListQuery struct {
	ContextID int64
	Context   string
	FolderID  int64
	Folder    string
	GoalID    int64
	Goal      string
	Priority  string `validate:"omitempty,oneof=Top top High high Medium medium Low low Negative negative"`
	Status    string `validate:"omitempty,oneof=None NextAction Active Planning Delegated Waiting Hold Postponed Someday Canceled Reference none nextaction active planning delegated waiting hold postponed someday canceled reference"`

	DueDate      string `validate:"omitempty,datetime=2006-01-02" json:"due_date" description:"format 2021-01-01"`
	SubTasksMode string `validate:"omitempty,oneof=Inline Hidden Indented inline hidden indented"`

	Format     string `validate:"omitempty,oneof=name json yaml"`
	Incomplete bool   `validate:"omitempty"`
	Complete   bool   `validate:"omitempty"`
	// TODO
	// Tags
}

var cmdQ = &cmdListQuery{}

func (q *cmdListQuery) PrepareIDs(contextSvc services.ContextService, goalSvc services.GoalService,
	folderSvc services.FolderService) error {
	if q.ContextID == 0 && q.Context != "" {
		// TODO case sensitive
		c, err := contextSvc.Find(q.Context)
		if err != nil {
			return errors.Wrap(err, "get context by name")
		}
		q.ContextID = c.ID
	}
	if q.FolderID == 0 && q.Folder != "" {
		// TODO case sensitive
		f, err := folderSvc.Find(q.Folder)
		if err != nil {
			return errors.Wrap(err, "get folder by name")
		}
		q.FolderID = f.ID
	}
	if q.GoalID == 0 && q.Goal != "" {
		// TODO case sensitive
		g, err := goalSvc.Find(q.Goal)
		if err != nil {
			return errors.Wrap(err, "get goal by name")
		}
		q.GoalID = g.ID
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
	// TODO validate incomplete and complete
	if q.Incomplete {
		in := true
		query.Incomplete = &in
	}
	// FIXME, not supported now, wait syncer support
	// remoteSvc.List
	if q.Complete {
		in := false
		query.Incomplete = &in
	}

	return query, nil
}

func NewListCmd(f *cmdutil.Factory) *cobra.Command {
	log := logrus.StandardLogger()
	cmd := &cobra.Command{
		Use:   "list",
		Args:  cobra.NoArgs,
		Short: "List tasks",
		Example: heredoc.Doc(`
			$ toodledo tasks list
			$ toodledo tasks list --context Work
			$ toodledo tasks list --context-id 4455
			$ toodledo tasks list --folder inbox
			$ toodledo tasks list --folder-id 4455
			$ toodledo tasks list --goal landing-moon
			$ toodledo tasks list --goal-id 4455
			$ toodledo tasks list --priority High
			$ toodledo tasks list --status Active
			$ toodledo tasks list --due-date "2020-01-01"
		`),
		Run: func(cmd *cobra.Command, args []string) {
			err := utils.FillQueryByFlags(cmd, cmdQ)
			if err != nil {
				log.WithError(err).Fatal("failed")
			}
			validate := validator.New()
			err = validate.Struct(cmdQ)
			if err != nil {
				log.WithError(err).Fatal("validate failed")
			}
			// FIXME cli is using TUI App, it contains local data
			app, err := injector.InitCLIApp()
			if err != nil {
				log.Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			appExt, _ := injector.InitTUIApp()

			svc := appExt.TaskLocalSvc
			contextSvc := app.ContextSvc
			folderSvc := app.FolderSvc
			goalSvc := app.GoalSvc
			taskRichSvc := app.TaskRichSvc
			fetcher := fetchers.NewToodledoFetchFunc(
				log,
				appExt.FolderLocalSvc,
				appExt.ContextLocalSvc,
				appExt.GoalLocalSvc,
				appExt.TaskLocalSvc,
				app.AccountSvc,
			)
			err = fetcher.Fetch(fetchers.NewNoOpStatusDescriber())
			if err != nil {
				log.WithError(err).Fatal("fetch failed")
				return
			}

			err = cmdQ.PrepareIDs(contextSvc, goalSvc, folderSvc)
			if err != nil {
				log.WithError(err).Fatal("prepare ids failed")
				return
			}
			q, err := cmdQ.ToQuery()
			if err != nil {
				log.WithError(err).Fatal("parse query failed")
			}

			tasks, err := svc.ListAllByQuery(q)
			if err != nil {
				log.Error(err)
				return
			}
			sorted, err := services.SortSubTasks(tasks, subtasksview.ModeString2Type(cmdQ.SubTasksMode))
			rts, _ := taskRichSvc.RichThem(sorted)
			switch cmdQ.Format {
			case "json":
				bs, _ := json.Marshal(rts)
				fmt.Println(string(bs))
			case "yaml":
				bs, _ := yaml.Marshal(rts)
				fmt.Println(string(bs))
			default:
				fmt.Println(render.Tables4RichTasks(rts))
			}
		},
	}

	err := utils.BindFlagsByQuery(cmd, *cmdQ)
	if err != nil {
		logrus.WithError(err).Fatal("bind flags failed")
	}
	return cmd
}
