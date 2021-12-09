// Code generated by "stringer -type=TaskStatus"; DO NOT EDIT.

package taskstatus

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[None-0]
	_ = x[NextAction-1]
	_ = x[Active-2]
	_ = x[Planning-3]
	_ = x[Delegated-4]
	_ = x[Waiting-5]
	_ = x[Hold-6]
	_ = x[Postponed-7]
	_ = x[Someday-8]
	_ = x[Canceled-9]
	_ = x[Reference-10]
}

const _TaskStatus_name = "NoneNextActionActivePlanningDelegatedWaitingHoldPostponedSomedayCanceledReference"

var _TaskStatus_index = [...]uint8{0, 4, 14, 20, 28, 37, 44, 48, 57, 64, 72, 81}

// String ...
func (i TaskStatus) String() string {
	if i < 0 || i >= TaskStatus(len(_TaskStatus_index)-1) {
		return "TaskStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TaskStatus_name[_TaskStatus_index[i]:_TaskStatus_index[i+1]]
}
