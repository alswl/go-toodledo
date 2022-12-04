//go:generate stringer -type=DueDateMode

package tasks

type DueDateMode int64

const (
	DueDateModeDueBy      DueDateMode = 0
	DueDateModeDueOn      DueDateMode = 1
	DueDateModeDueAfter   DueDateMode = 2
	DueDateModeOptionally DueDateMode = 3
)

func DueDateModeAll() []DueDateMode {
	return []DueDateMode{DueDateModeDueBy, DueDateModeDueOn, DueDateModeDueAfter, DueDateModeOptionally}
}

func DueDateModeValue2Type(input int64) DueDateMode {
	for _, x := range DueDateModeAll() {
		if x == DueDateMode(input) {
			return x
		}
	}
	return DueDateModeDueBy
}
