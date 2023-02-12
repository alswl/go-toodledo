// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	ccontext "context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TaskEdit task edit
//
// swagger:model TaskEdit
type TaskEdit struct {

	// A GMT unix timestamp for when the task was last modified.
	Added *int64 `json:"added,omitempty"`

	// The user id of the collaborator who assigned the task (Subscription required for user).
	Addedby *int64 `json:"addedby,omitempty"`

	// An array of attachment items. Each item will contain the following three fields. Attachments are currently read only. You can use the id number to reference the outline, list or note that you can get via this API. File's are not currently readable, but we plan to add this functionality soon.
	// id: The unique id number for the attachment
	// kind: The kind of attachment (file,note,outline,list)
	// name: The display name of the attachment
	//
	Attatchment string `json:"attatchment,omitempty"`

	// This is used with Subscriptions that have access to subtasks. This will indicate the number of child tasks that this task has. This will be 0 for subtasks or for regular tasks without subtasks.
	Children *int64 `json:"children,omitempty"`

	// A GMT unix timestamp for when the task was completed. If the task is not completed, the value will be 0. Toodledo does not track the time that a task was completed, so tasks will always appear to be completed at noon.
	Completed *int64 `json:"completed,omitempty"`

	// The id number of the context. Omit this field or set it to 0 to leave the task unassigned to a context.
	Context *int64 `json:"context,omitempty"`

	// A GMT unix timestamp for when the task is due. The time component of this timestamp doesn't matter. When fetching from the server, it will always be noon.
	Duedate *int64 `json:"duedate,omitempty"`

	// An integer representing the due date modifier.
	// 0 = Due By
	// 1 = Due On (=)
	// 2 = Due After (>)
	// 3 = Optionally (?)
	//
	Duedatemod *int64 `json:"duedatemod,omitempty"`

	// A GMT unix timestamp for when the task is due. If the task does not have a time set, then this will be 0. If the task has a duetime without a duedate set, then the date component of this timestamp will be Jan 1, 1970. Times are stored as floating times. In other words, 10am is always 10am, regardless of your timezone. You can convert this timestamp to a GMT string and display the time component without worrying about timezones.
	Duetime *int64 `json:"duetime,omitempty"`

	// The id number of the folder. Omit this field or set it to 0 to leave the task unassigned to a folder.
	Folder int64 `json:"folder,omitempty"`

	// The id number of the goal. Omit this field or set it to 0 to leave the task unassigned to a goal.
	Goal *int64 `json:"goal,omitempty"`

	// The server id number for this task. It is guaranteed to be unique per account, but two different accounts may have two different tasks with the same id number.
	ID int64 `json:"id,omitempty"`

	// An integer representing the number of minutes that the task will take to complete.
	Length *int64 `json:"length,omitempty"`

	// The id number of the location. Omit this field or set it to 0 to leave the task unassigned to a location.
	Location *int64 `json:"location,omitempty"`

	// A text string up to 1024 bytes long for storing metadata about the task. This is useful for syncing data that cannot otherwise be synced to Toodledo. This data is unique per task ID. This data is private to your App ID. Neither the user, nor other App IDs can see the data that you put in here. Because of an implementation detail, using the meta field introduces extra latency to each API call, so you should only use this field when necessary.
	Meta *string `json:"meta,omitempty"`

	// A GMT unix timestamp for when the task was last modified.
	Modified *int64 `json:"modified,omitempty"`

	// A text string up to 32,000 bytes long. New lines should be sent as \n.
	Note *string `json:"note,omitempty"`

	// This is used with Subscriptions that have access to subtasks. This is an integer that indicates the manual order of subtasks within the parent task. Currently this is read-only.
	Order *int64 `json:"order,omitempty"`

	// This is used with Subscriptions that have access to subtasks. To create a subtask, set this to the id number of the parent task. The default is 0, which creates a normal task.
	Parent *int64 `json:"parent,omitempty"`

	// If the task was repeated from another task, this will contain the id number of the previous version of this task.
	Previous *int64 `json:"previous,omitempty"`

	// An integer that represents the priority.
	// -1 = Negative
	// 0 = Low
	// 1 = Medium
	// 2 = High
	// 3 = Top
	//
	Priority *int64 `json:"priority,omitempty"`

	// A read-only field that contains the id number of the task that this task is a reference to. This is useful for tasks that are created from another task.
	Ref *string `json:"ref,omitempty"`

	// An integer that represents the number of minutes prior to the duedate/time that a reminder will be sent. Set it to 0 for no reminder. Values will be constrained to this list of valid numbers (0, 1, 15, 30, 45, 60, 90, 120, 180, 240, 1440, 2880, 4320, 5760, 7200, 8640, 10080, 20160, 43200). Additionally, if the user does not have a Subscription, the only valid numbers are 0,60. If you submit an invalid number, it will be rounded up or down to a valid non zero value.
	Remind *int64 `json:"remind,omitempty"`

	// A string indicating how the task repeats. When a task is rescheduled, it is moved forward to the new date. For record keeping purposes, a completed copy of this task will be added to the user's list. It will have a new ID number and will be already completed. To make a task no longer repeat, set this field to an empty string.
	//
	// This string is in the standard iCal RRULE format. For example: "FREQ=DAILY" or "FREQ=WEEKLY;BYDAY=TU". Not every valid iCal RRULE is understood by Toodledo, but we will be improving our ability to understand more options in the future. Please read our repeat format FAQ for details about how we can currently repeat tasks. Note that users enter their repeat settings using a GUI or by typing a human readable string. These get converted to iCal rules behind the scenes.
	//
	// There are three enhancements to the iCal RRULE format that we have made to support some advanced Toodledo features.
	// Subtasks can repeat based on their parent's repeat value. There is not a comparable iCal RRULE for this, so we have a custom RRULE of "PARENT" to indicate this setting.
	// Tasks can be set to repeat from their due-date or their completion date. There is not a way to indicate this in a standard iCAL RRULE. To indicate this, we have a custom string of ";FROMCOMP" that we append to the RRULE if the task is set to repeat from the completion date. The absence of this string means that the task repeats from the due-date.
	// Normally, when a task is rescheduled it moves forwards by 1 occurrence. If the user has procrastinated, the new due-date could still be in the past. Toodledo will have the option to indicate that certain repeating tasks should be rescheduled to the next future occurance of the task. If this is the case, the custom ";FASTFORWARD" string will be appended to the RRULE.
	//
	Repeat *string `json:"repeat,omitempty"`

	// The user id of the person who owns the task that is being shared with the current user. If the current user is not the owner, then they cannot make changes to the collaboration settings for this task, although they can make other changes. Read only.
	Sharedowner *int64 `json:"sharedowner,omitempty"`

	// An array of user ids for people that this task is shared with, other than myself and the owner. Read only.
	Sharedwith *int64 `json:"sharedwith,omitempty"`

	// A boolean (0 or 1) that indicates if the task has a star or not.
	Star int64 `json:"star,omitempty"`

	// A GMT unix timestamp for when the task starts. The time component of this timestamp will always be noon.
	Startdate *int64 `json:"startdate,omitempty"`

	// A GMT unix timestamp for when the task starts. If the task does not have a time set, then this will be 0. If the task has a starttime without a startdate set, then the date component of this timestamp will be Jan 1, 1970. Times are stored as floating times. In other words, 10am is always 10am, regardless of your timezone. You can convert this timestamp to a GMT string and display the time component without worrying about timezones.
	Starttime *int64 `json:"starttime,omitempty"`

	// An integer that represents the status of the task.
	// 0 = None
	// 1 = Next Action
	// 2 = Active
	// 3 = Planning
	// 4 = Delegated
	// 5 = Waiting
	// 6 = Hold
	// 7 = Postponed
	// 8 = Someday
	// 9 = Canceled
	// 10 = Reference
	//
	Status *int64 `json:"status,omitempty"`

	// A comma separated string listing the tags assigned to this task. Up to 250 characters.
	Tag *string `json:"tag,omitempty"`

	// The value in the timer field indicates the number of seconds that have elapsed for the timer not including the current session.
	Timer *int64 `json:"timer,omitempty"`

	// If the timer is currently on, this will contain a unix timestamp indicating the last time that the timer was started. Therefore, if the timer is currently on, you will need to calculate the elapsed time when you present it to the user. This calculation is: Total Time=timer+(now-timeron). Where "now" is a unix timestamp for the current time.
	Timeron *int64 `json:"timeron,omitempty"`

	// A string for the name of the task. Up to 255 characters.
	Title *string `json:"title,omitempty"`

	// A read-only field that indicates how the task was added. These are the possible values:
	// 0: Main website
	// 1: Email Import
	// 2: Firefox Addon
	// 3: This API
	// 4: Widgets (Google Gadget, etc)
	// 5: Not used
	// 6: Mobile Phone
	// 7: iPhone App
	// 8: Import Tools
	// 9: Twitter
	//
	Via string `json:"via,omitempty"`
}

// Validate validates this task edit
func (m *TaskEdit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this task edit based on context it is used
func (m *TaskEdit) ContextValidate(ctx ccontext.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskEdit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskEdit) UnmarshalBinary(b []byte) error {
	var res TaskEdit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
