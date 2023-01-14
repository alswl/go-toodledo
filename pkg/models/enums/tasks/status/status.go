//go:generate stringer -type=Status
package status

type Status int

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

func All() []Status {
	return []Status{
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
}

var mapping = map[string]Status{
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

func Value2Type(input int64) Status {
	for _, x := range All() {
		if x == Status(input) {
			return x
		}
	}
	return None
}

func String2Type(input string) Status {
	for k, v := range mapping {
		if k == input {
			return v
		}
	}
	return None
}
