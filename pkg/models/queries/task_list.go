package queries

import (
	"github.com/alswl/go-toodledo/pkg/models/enums/tasks/priority"
	"github.com/alswl/go-toodledo/pkg/models/enums/tasks/status"
)

type TaskListQuery struct {
	Title string
	// FIXME how to present no content?
	ContextID int64
	FolderID  int64
	GoalID    int64
	DueDate   string
	// Priority, low is zero, is default value, Priority should be pointer
	Priority   *priority.Priority
	Status     *status.Status
	Incomplete *bool
}
