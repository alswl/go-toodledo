package enums

import "strings"

// TaskField ...
type TaskField string

// TaskFieldID ...
const (
	TaskFieldID              TaskField = "id"
	TaskFieldTitle           TaskField = "title"
	TaskFieldModified        TaskField = "modified"
	TaskFieldCompleted       TaskField = "completed"
	TaskFieldFolder          TaskField = "folder"
	TaskFieldContext         TaskField = "context"
	TaskFieldGoal            TaskField = "goal"
	TaskFieldLocation        TaskField = "location"
	TaskFieldTag             TaskField = "tag"
	TaskFieldStartDate       TaskField = "startdate"
	TaskFieldDueDate         TaskField = "duedate"
	TaskFieldDueDateModified TaskField = "duedatemod"
	TaskFieldStartTime       TaskField = "starttime"
	TaskFieldDueTime         TaskField = "duetime"
	TaskFieldRemind          TaskField = "remind"
	TaskFieldRepeat          TaskField = "repeat"
	TaskFieldStatus          TaskField = "status"
	TaskFieldStar            TaskField = "star"
	TaskFieldPriority        TaskField = "priority"
	TaskFieldLength          TaskField = "length"
	TaskFieldTimer           TaskField = "timer"
	TaskFieldAdded           TaskField = "added"
	TaskFieldNote            TaskField = "note"
	TaskFieldParent          TaskField = "parent"
	TaskFieldChildren        TaskField = "children"
	TaskFieldOrder           TaskField = "order"
	TaskFieldMeta            TaskField = "meta"
	TaskFieldPrevious        TaskField = "previous"
	TaskFieldAttachment      TaskField = "attachment"
	TaskFieldShared          TaskField = "shared"
	TaskFieldAddedBy         TaskField = "addedby"
	TaskFieldVia             TaskField = "via"
	TaskFieldAttachments     TaskField = "attachments"
)

// DefaultTaskFields show fields always return, id, title, modified, completed.
func DefaultTaskFields() []TaskField {
	return []TaskField{
		TaskFieldID,
		TaskFieldTitle,
		TaskFieldCompleted,
		TaskFieldModified,
	}
}

func TinyTaskFields() []TaskField {
	return []TaskField{
		TaskFieldFolder,
		TaskFieldContext,
		TaskFieldGoal,
		TaskFieldDueDate,
		TaskFieldDueTime,
		TaskFieldRepeat,
		TaskFieldStatus,
		TaskFieldStar,
		TaskFieldPriority,
		TaskFieldParent,
		TaskFieldChildren,
	}
}

func GeneralTaskFields() []TaskField {
	return append(TinyTaskFields(),
		TaskFieldLocation,
		TaskFieldTag,
		TaskFieldStartDate,
		TaskFieldDueDateModified,
		TaskFieldStartTime,
		TaskFieldRemind,
		TaskFieldLength,
		TaskFieldTimer,
		TaskFieldAdded,
		TaskFieldNote,
		TaskFieldOrder,
		TaskFieldMeta,
		TaskFieldPrevious,
		TaskFieldAttachment,
		TaskFieldShared,
		TaskFieldAddedBy,
		TaskFieldVia,
	)
}

func FullTaskFields() []TaskField {
	return append(GeneralTaskFields(),
		TaskFieldAttachments,
	)
}

// TaskFields2String ...
func TaskFields2String(fields []TaskField) string {
	var fs []string
	for _, f := range fields {
		fs = append(fs, string(f))
	}
	return strings.Join(fs, ",")
}
