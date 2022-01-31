//go:generate stringer -type=Status

package tasks

// Status ...
type Status int

// StatusNone ...
const (
	StatusNone       Status = 0
	StatusNextAction Status = 1
	StatusActive     Status = 2
	StatusPlanning   Status = 3
	StatusDelegated  Status = 4
	StatusWaiting    Status = 5
	StatusHold       Status = 6
	StatusPostponed  Status = 7
	StatusSomeday    Status = 8
	StatusCanceled   Status = 9
	StatusReference  Status = 10
)

// StatusAll ...
var StatusAll = []Status{
	StatusNone,
	StatusNextAction,
	StatusActive,
	StatusPlanning,
	StatusDelegated,
	StatusWaiting,
	StatusHold,
	StatusPostponed,
	StatusSomeday,
	StatusCanceled,
	StatusReference,
}

var StatusMap = map[string]Status{
	"none":        StatusNone,
	"next_action": StatusNextAction,
	"nextaction":  StatusNextAction,
	"active":      StatusActive,
	"planning":    StatusPlanning,
	"delegated":   StatusDelegated,
	"waiting":     StatusWaiting,
	"hold":        StatusHold,
	"postponed":   StatusPostponed,
	"someday":     StatusSomeday,
	"canceled":    StatusCanceled,
	"reference":   StatusReference,
}

// StatusValue2Type ...
func StatusValue2Type(input int64) Status {
	for _, x := range StatusAll {
		if x == Status(input) {
			return x
		}
	}
	return StatusNone
}

func StatusString2Type(input string) Status {
	for k, v := range StatusMap {
		if k == input {
			return v
		}
	}
	return StatusNone
}
