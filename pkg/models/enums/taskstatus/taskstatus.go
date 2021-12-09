//go:generate stringer -type=TaskStatus

package taskstatus

// TaskStatus ...
type TaskStatus int

// None ...
const (
	None       TaskStatus = 0
	NextAction TaskStatus = 1
	Active     TaskStatus = 2
	Planning   TaskStatus = 3
	Delegated  TaskStatus = 4
	Waiting    TaskStatus = 5
	Hold       TaskStatus = 6
	Postponed  TaskStatus = 7
	Someday    TaskStatus = 8
	Canceled   TaskStatus = 9
	Reference  TaskStatus = 10
)

// All ...
var All = []TaskStatus{
	None,
	NextAction,
	Active,
	Planning,
	Delegated,
	Waiting,
	Hold,
	Postponed,
	Someday,
	Canceled,
	Reference,
}

// Value2Type ...
func Value2Type(input int64) TaskStatus {
	for _, x := range All {
		if x == TaskStatus(input) {
			return x
		}
	}
	return None
}
