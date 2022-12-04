package queries

import (
	"fmt"

	"github.com/alswl/go-toodledo/pkg/models/enums/tasks/priority"
	"github.com/alswl/go-toodledo/pkg/models/enums/tasks/status"
)

type TaskListQuery struct {
	Title string
	// ContextID, 0 for all, -1 for none
	ContextID int64
	// FolderID, 0 for all, -1 for none
	FolderID int64
	// GoalID, 0 for all, -1 for none
	GoalID  int64
	DueDate string
	// Priority, low is zero, is default value, Priority should be pointer
	Priority   *priority.Priority
	Status     *status.Status
	Incomplete *bool
}

func (q TaskListQuery) UniqString() string {
	return fmt.Sprintf(
		"%s-%d-%d-%d-%s-%s-%s-%v",
		q.Title,
		q.ContextID,
		q.FolderID,
		q.GoalID,
		q.DueDate,
		q.Priority,
		q.Status,
		q.Incomplete,
	)
}
