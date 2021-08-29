package enums

import "strings"

type TaskField string

var (
	TASK_FIELD_ID                TaskField = "id"
	TASK_FIELD_TITLE             TaskField = "title"
	TASK_FIELD_MODIFIED          TaskField = "modified"
	TASK_FIELD_COMPLETED         TaskField = "completed"
	TASK_FIELD_FOLDER            TaskField = "folder"
	TASK_FIELD_CONTEXT           TaskField = "context"
	TASK_FIELD_GOAL              TaskField = "goal"
	TASK_FIELD_LOCATION          TaskField = "location"
	TASK_FIELD_TAG               TaskField = "tag"
	TASK_FIELD_START_DATE        TaskField = "startdate"
	TASK_FIELD_DUE_DATE          TaskField = "duedate"
	TASK_FIELD_DUE_DATE_MODIFIED TaskField = "duedatemod"
	TASK_FIELD_START_TIME        TaskField = "starttime"
	TASK_FIELD_DUE_TIME          TaskField = "duetime"
	TASK_FIELD_REMIND            TaskField = "remind"
	TASK_FIELD_REPEAT            TaskField = "repeat"
	TASK_FIELD_STATUS            TaskField = "status"
	TASK_FIELD_STAR              TaskField = "star"
	TASK_FIELD_PRIORITY          TaskField = "priority"
	TASK_FIELD_LENGTH            TaskField = "length"
	TASK_FIELD_TIMER             TaskField = "timer"
	TASK_FIELD_ADDED             TaskField = "added"
	TASK_FIELD_NOTE              TaskField = "note"
	TASK_FIELD_PARENT            TaskField = "parent"
	TASK_FIELD_CHILDREN          TaskField = "children"
	TASK_FIELD_ORDER             TaskField = "order"
	TASK_FIELD_META              TaskField = "meta"
	TASK_FIELD_PREVIOUS          TaskField = "previous"
	TASK_FIELD_ATTACHMENT        TaskField = "attachment"
	TASK_FIELD_SHARED            TaskField = "shared"
	TASK_FIELD_ADDED_BY          TaskField = "addedby"
	TASK_FIELD_VIA               TaskField = "via"
	TASK_FIELD_ATTACHMENTS       TaskField = "attachments"
)

// DefaultTaskFields show fields always return, id, title, modified, completed
var DefaultTaskFields = []TaskField{
	TASK_FIELD_ID,
	TASK_FIELD_TITLE,
	TASK_FIELD_COMPLETED,
	TASK_FIELD_MODIFIED,
}

var TinyTaskFields = []TaskField{
	TASK_FIELD_FOLDER,
	TASK_FIELD_CONTEXT,
	TASK_FIELD_GOAL,
	TASK_FIELD_DUE_DATE,
	TASK_FIELD_DUE_TIME,
	TASK_FIELD_REPEAT,
	TASK_FIELD_STATUS,
	TASK_FIELD_STAR,
	TASK_FIELD_PRIORITY,
	TASK_FIELD_PARENT,
	TASK_FIELD_CHILDREN,
}

var GeneralTaskFields = append(TinyTaskFields,
	TASK_FIELD_LOCATION,
	TASK_FIELD_TAG,
	TASK_FIELD_START_DATE,
	TASK_FIELD_DUE_DATE_MODIFIED,
	TASK_FIELD_START_TIME,
	TASK_FIELD_REMIND,
	TASK_FIELD_LENGTH,
	TASK_FIELD_TIMER,
	TASK_FIELD_ADDED,
	TASK_FIELD_NOTE,
	TASK_FIELD_ORDER,
	TASK_FIELD_META,
	TASK_FIELD_PREVIOUS,
	TASK_FIELD_ATTACHMENT,
	TASK_FIELD_SHARED,
	TASK_FIELD_ADDED_BY,
	TASK_FIELD_VIA,
)

var FullTaskFields = append(GeneralTaskFields,
	TASK_FIELD_ATTACHMENTS,
)

func TaskFields2String(fields []TaskField) string {
	var fs []string
	for _, f := range fields {
		fs = append(fs, string(f))
	}
	return strings.Join(fs, ",")
}
