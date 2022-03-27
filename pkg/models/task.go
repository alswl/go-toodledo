package models

import "time"

type RichTask struct {
	Task
	TheContext Context `json:"the_context"`
	TheFolder  Folder  `json:"the_folder"`
	TheGoal    Goal    `json:"the_goal"`
	//AddedByUser *Account `json:"added_by_user"`
	//ParentTask *Task `json:"parent_task"`
	//PreviousTask *Task `json:"previous_task"`
}

func (t RichTask) DueDate() time.Time {
	if t.Duedate == 0 {
		return time.Time{}
	}
	// TODO get timezone from toodledo
	return time.Unix(t.Duedate, 0).In(time.Local)
}

func (t RichTask) DueTime() time.Time {
	if t.Duetime == 0 {
		return time.Time{}
	}
	// due time is calculated in UTC(ignore timezone)
	return time.Unix(t.Duetime, 0).In(time.UTC)
}

func (t RichTask) Due() string {
	var output = ""
	if !t.DueDate().IsZero() {
		output += t.DueDate().Format("2006-01-02")
	}
	if !t.DueTime().IsZero() {
		if output != "" {
			output += " "
		}
		output += t.DueTime().Format("15:04")
	}
	return output
}
