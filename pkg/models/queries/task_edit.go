package queries

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"strings"
	"time"
)

// TaskEditQuery is edit query model of Task
type TaskEditQuery struct {
	TaskWritePartialQuery

	ID    int64  `description:"" validate:"required"`
	Title string `description:""`
}

// ToModel converts TaskCreateQuery to Task
func (q *TaskEditQuery) ToModel() *models.Task {
	t := &models.Task{
		ID:      q.ID,
		Title:   q.Title,
		Context: q.ContextID,
		Folder:  q.FolderID,
		Goal:    q.GoalID,

		Duedatemod: int64(q.DueDateMode),

		Length: q.Length,
		//Location:     q.Location,
		Note:     q.Note,
		Parent:   q.Parent,
		Priority: int64(q.Priority),
		//Remind:     q.Remind,
		//Repeat:     q.Repeat,
		Status: int64(q.Status),
		Timer:  q.Timer,
		//Via:        q.Via,
	}
	if q.DueDate != "" {
		dueDate, err := time.Parse("2006-01-02", q.DueDate)
		if err != nil {
			return nil
		}
		t.Duedate = dueDate.Unix()
	}
	if q.DueTime != 0 {
		t.Duetime = q.DueTime
	}
	if q.Star {
		t.Star = 1
	}
	if len(q.Tag) > 0 {
		t.Tag = strings.Join(q.Tag, ",")
	}
	if !q.TimerOne.IsZero() {
		t.Timerone = q.TimerOne.Unix()
	}

	return t
}
