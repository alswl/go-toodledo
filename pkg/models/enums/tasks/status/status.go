//go:generate stringer -type=Status
package status

// Status ...
type Status int

// StatusNone ...
const (
	None       Status = 0
	NextAction Status = 1
	Active     Status = 2
	Planning   Status = 3
	Delegated  Status = 4
	Waiting    Status = 5
	Hold       Status = 6
	Postponed  Status = 7
	Someday    Status = 8
	Canceled   Status = 9
	Reference  Status = 10
)

// StatusAll ...
var StatusAll = []Status{
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

var StatusMap = map[string]Status{
	"none":        None,
	"next_action": NextAction,
	"nextaction":  NextAction,
	"active":      Active,
	"planning":    Planning,
	"delegated":   Delegated,
	"waiting":     Waiting,
	"hold":        Hold,
	"postponed":   Postponed,
	"someday":     Someday,
	"canceled":    Canceled,
	"reference":   Reference,
}

// StatusValue2Type ...
func StatusValue2Type(input int64) Status {
	for _, x := range StatusAll {
		if x == Status(input) {
			return x
		}
	}
	return None
}

func StatusString2Type(input string) Status {
	for k, v := range StatusMap {
		if k == input {
			return v
		}
	}
	return None
}
