//go:generate stringer -type=Priority

package tasks

// Priority ...
type Priority int

const (
	PriorityNegative Priority = -1
	PriorityLow      Priority = 0
	PriorityMedium   Priority = 1
	PriorityHigh     Priority = 2
	PriorityTop      Priority = 3
)

// PriorityAll ...
var PriorityAll = []Priority{
	PriorityNegative,
	PriorityLow,
	PriorityMedium,
	PriorityHigh,
	PriorityTop,
}

// PriorityValue2Type ...
func PriorityValue2Type(input int64) Priority {
	for _, x := range PriorityAll {
		if x == Priority(input) {
			return x
		}
	}
	return PriorityMedium
}
