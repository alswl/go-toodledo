package queries

import (
	"github.com/alswl/go-toodledo/pkg/models"
	"github.com/alswl/go-toodledo/pkg/models/enums/tasks"
	"strings"
	"time"
)

type TaskSearchQuery struct {
	Title     string
	ContextID int64
	FolderID  int64
	GoalID    int64
	DueDate   string
	// Priority, low is zero, is default value, Priority should be pointer
	Priority *tasks.Priority
	// XXX add Status
	// Status
}

// TaskCreateQuery is query model of Task
type TaskCreateQuery struct {
	// TODO fields missing
	Title     string `description:"" validate:"required"`
	ContextID int64
	FolderID  int64
	GoalID    int64

	// TODO fields using go type
	DueDate     string
	DueDateMode tasks.DueDateMode
	// TODO fields using go type
	DueTime int64

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

// ToModel converts TaskCreateQuery to Task
func (q *TaskCreateQuery) ToModel() *models.Task {
	t := &models.Task{
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

// TaskCreateQueryBuilder is the builder of TaskCreateQuery
type TaskCreateQueryBuilder struct {
	query TaskCreateQuery
}

func NewTaskCreateQueryBuilder() *TaskCreateQueryBuilder {
	return &TaskCreateQueryBuilder{}
}

// WithTitle sets Title
func (b *TaskCreateQueryBuilder) WithTitle(title string) *TaskCreateQueryBuilder {
	b.query.Title = title
	return b
}

// WithContextID sets ContextID
func (b *TaskCreateQueryBuilder) WithContextID(contextID int64) *TaskCreateQueryBuilder {
	b.query.ContextID = contextID
	return b
}

// WithFolderID sets FolderID
func (b *TaskCreateQueryBuilder) WithFolderID(folderID int64) *TaskCreateQueryBuilder {
	b.query.FolderID = folderID
	return b
}

// WithGoalID sets GoalID
func (b *TaskCreateQueryBuilder) WithGoalID(goalID int64) *TaskCreateQueryBuilder {
	b.query.GoalID = goalID
	return b
}

// WithDueDate sets DueDate
func (b *TaskCreateQueryBuilder) WithDueDate(dueDate string) *TaskCreateQueryBuilder {
	b.query.DueDate = dueDate
	return b
}

// WithDueTime sets DueTime
func (b *TaskCreateQueryBuilder) WithDueTime(dueTime int64) *TaskCreateQueryBuilder {
	b.query.DueTime = dueTime
	return b
}

// WithDueDateMode sets DueDateMode
func (b *TaskCreateQueryBuilder) WithDueDateMode(dueDateMode tasks.DueDateMode) *TaskCreateQueryBuilder {
	b.query.DueDateMode = dueDateMode
	return b
}

// WithLength sets Length
func (b *TaskCreateQueryBuilder) WithLength(length int64) *TaskCreateQueryBuilder {
	b.query.Length = length
	return b
}

// WithNote sets Note
func (b *TaskCreateQueryBuilder) WithNote(note string) *TaskCreateQueryBuilder {
	b.query.Note = note
	return b
}

// WithParent sets Parent
func (b *TaskCreateQueryBuilder) WithParent(parent int64) *TaskCreateQueryBuilder {
	b.query.Parent = parent
	return b
}

// WithPriority sets Priority
func (b *TaskCreateQueryBuilder) WithPriority(priority tasks.Priority) *TaskCreateQueryBuilder {
	b.query.Priority = priority
	return b
}

// WithStar sets Star
func (b *TaskCreateQueryBuilder) WithStar(star bool) *TaskCreateQueryBuilder {
	b.query.Star = star
	return b
}

// WithStatus sets Status
func (b *TaskCreateQueryBuilder) WithStatus(status tasks.Status) *TaskCreateQueryBuilder {
	b.query.Status = status
	return b
}

// WithTag sets Tag
func (b *TaskCreateQueryBuilder) WithTag(tag string) *TaskCreateQueryBuilder {
	b.query.Tag = append(b.query.Tag, tag)
	return b
}

// WithTimer sets Timer
func (b *TaskCreateQueryBuilder) WithTimer(timer int64) *TaskCreateQueryBuilder {
	b.query.Timer = timer
	return b
}

// WithTimerOne sets TimeOne
func (b *TaskCreateQueryBuilder) WithTimerOne(timerOne time.Time) *TaskCreateQueryBuilder {
	b.query.TimerOne = timerOne
	return b
}

// Build returns TaskCreateQuery
func (b *TaskCreateQueryBuilder) Build() *TaskCreateQuery {
	return &b.query
}

// TODO
type TaskEditQuery struct {
	TaskCreateQuery
}
