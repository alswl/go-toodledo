package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/teambition/rrule-go"

	"github.com/alswl/go-toodledo/pkg/models/enums/tasks/priority"
	"github.com/alswl/go-toodledo/pkg/models/enums/tasks/status"

	"github.com/alswl/go-toodledo/pkg/utils"
	utilsrrule "github.com/alswl/go-toodledo/pkg/utils/rrule"
	utilstime "github.com/alswl/go-toodledo/pkg/utils/time"
)

type RichTask struct {
	Task
	TheContext *Context `json:"the_context"`
	TheFolder  *Folder  `json:"the_folder"`
	TheGoal    *Goal    `json:"the_goal"`
	// AddedByUser *Account `json:"added_by_user"`
	// ParentTask *Task `json:"parent_task"`
	// PreviousTask *Task `json:"previous_task"`

	// Due
	// Repeat
	// Tag []string
}

func (t RichTask) TheDueDate() time.Time {
	if t.Duedate == 0 {
		return time.Time{}
	}
	// TODO get timezone from toodledo
	return time.Unix(t.Duedate, 0).In(utils.DefaultTimeZone)
}

func (t RichTask) TheDueTime() time.Time {
	if t.Duetime == 0 {
		return time.Time{}
	}
	// due time is calculated in UTC(ignore timezone)
	return time.Unix(t.Duetime, 0).In(time.UTC)
}

func (t RichTask) DueString() string {
	var output = ""
	if !t.TheDueDate().IsZero() {
		output += t.TheDueDate().Format("2006-01-02")
	}
	if !t.TheDueTime().IsZero() {
		if output != "" {
			output += " "
		}
		output += t.TheDueTime().Format("15:04")
	}
	return output
}

func (t RichTask) RepeatString() string {
	if t.Task.Repeat == "" {
		return ""
	}
	r, err := rrule.StrToRRule(t.Task.Repeat)
	if err != nil {
		// TODO log
		return "."
	}
	return utilsrrule.ParseToodledoRRule(*r)
}

// TimerString return the timer string of the task
// if timer is on, * will be added in the front.
func (t RichTask) TimerString() string {
	if t.Timer == 0 && t.Timeron == 0 {
		return ""
	}
	// nolint:gomnd
	d := time.Duration(t.Timer) * time.Second
	now := time.Now()
	if t.Timeron != 0 {
		// TODO using user time zone
		timerOn := time.Unix(t.Timeron, 0).In(utils.DefaultTimeZone)
		d += now.Sub(timerOn)
	}
	readableDuration := utilstime.ParseDurationToReadable(d)
	if t.Timeron == 0 {
		return readableDuration
	}

	return fmt.Sprintf("*%s", readableDuration)
}

func (t RichTask) LengthString() string {
	if t.Length == 0 {
		return ""
	}
	d := time.Duration(t.Length) * time.Minute
	return utilstime.ParseDurationToReadable(d)
}

func (t RichTask) PriorityString() string {
	return priority.Value2Type(t.Priority).String()
}

func (t RichTask) TagString() string {
	return t.Tag
}

func (t RichTask) CompletedString() string {
	if t.Task.Completed == 0 {
		return "[ ]"
	}
	return "[X]"
}

// ThatContext returns the context of the task
// if the context is not set, it returns a default context.
func (t RichTask) ThatContext() Context {
	if t.TheContext == nil {
		return Context{}
	}
	return *t.TheContext
}

// ThatFolder returns the folder of the task
// if the folder is not set, it returns a default folder.
func (t RichTask) ThatFolder() Folder {
	if t.TheFolder == nil {
		return Folder{}
	}
	return *t.TheFolder
}

// ThatGoal returns the goal of the task
// if the goal is not set, it returns a default goal.
func (t RichTask) ThatGoal() Goal {
	if t.TheGoal == nil {
		return Goal{}
	}
	return *t.TheGoal
}

func (t RichTask) AddedString() string {
	// TODO use user time zone
	return time.Unix(t.Added, 0).In(utils.DefaultTimeZone).Format("2006-01-02 15:04:05")
}

func (t RichTask) StarString() string {
	return strconv.FormatBool(t.Star == 1)
}

func (t RichTask) StatusString() string {
	return status.Value2Type(t.Status).String()
}

func (t RichTask) ModifiedString() string {
	// TODO use user time zone
	return time.Unix(t.Modified, 0).In(utils.DefaultTimeZone).Format("2006-01-02 15:04:05")
}

func (t RichTask) Link() string {
	// TODO endpoint configurable
	return fmt.Sprintf("https://www.toodledo.com/tasks/index.php?#task_%d", t.ID)
}
