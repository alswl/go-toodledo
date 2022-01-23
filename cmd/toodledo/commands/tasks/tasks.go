package tasks

import (
	"github.com/alswl/go-toodledo/pkg/models/enums/tasks"
	"github.com/alswl/go-toodledo/pkg/models/queries"
	"github.com/spf13/cobra"
)

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage toodledo tasks",
	Args:  cobra.NoArgs,
}

// cmdQuery present the parameters for the command
// parse query with cmd
type cmdQuery struct {
	// TODO name
	ContextID int64
	// TODO name
	FolderID int64
	// TODO name
	GoalID   int64
	Priority string `validate:"omitempty,oneof=Top top High high Medium medium Low low Negative negative"`

	DueDate string `validate:"datetime=2006-01-02" json:"due_date" description:"format 2021-01-01"`
	// TODO
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

func init() {
	TaskCmd.AddCommand(listCmd, viewCmd, createCmd, deleteCmd, editCmd, completeCmd, uncompleteCmd)
}
