package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/alswl/go-toodledo/pkg/utils"
)

type RichTask struct {
	Task
	TheContext Context `json:"the_context"`
	TheFolder  Folder  `json:"the_folder"`
	TheGoal    Goal    `json:"the_goal"`
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
	return time.Unix(t.Duedate, 0).In(utils.ChinaTimeZone)
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
	return t.Task.Repeat
}

func (t RichTask) TimerString() string {
	if t.Timer == 0 && t.Timeron == 0 {
		return ""
	}
	if t.Timeron == 0 {
		duration := strconv.FormatInt((t.Timer)*1000*1000*1000, 10)
		return duration
	}
	now := time.Now().Unix()
	duration := strconv.FormatInt((now-t.Timeron+t.Timer)*1000*1000*1000, 10)
	return fmt.Sprintf("> %s", duration)
}

func (t RichTask) LengthString() string {
	if t.Length == 0 {
		return ""
	}
	return strconv.FormatInt(t.Length*1000*1000*1000, 10)
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
