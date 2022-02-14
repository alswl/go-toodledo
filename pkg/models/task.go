package models

type RichTask struct {
	*Task
	TheContext *Context `json:"the_context"`
	TheFolder  *Folder  `json:"the_folder"`
	TheGoal    *Goal    `json:"the_goal"`
	//AddedByUser *Account `json:"added_by_user"`
	//ParentTask *Task `json:"parent_task"`
	//PreviousTask *Task `json:"previous_task"`
}
