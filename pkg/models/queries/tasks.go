package queries

import (
	"github.com/alswl/go-toodledo/pkg/models/enums/tasks"
	"time"
)

// TaskCreateQuery is query model of Task
type TaskCreateQuery struct {
	Title   string
	Context string
	Folder  string
	Goal    string

	DueDate       time.Time
	DueDateMode   tasks.DueDateMode
	DueTimeHour   int64
	DueTimeMinute int64

	Length int64
	//Location     int64
	Note     string
	Parent   int64
	Priority tasks.Priority
	//Remind
	//Repeat iCal RRULE
	Star     bool
	Status   tasks.Status
	Tag      []string
	Timer    int64
	TimerOne time.Time
	//Via string
}

type TaskEditQuery struct {
	TaskCreateQuery
}
