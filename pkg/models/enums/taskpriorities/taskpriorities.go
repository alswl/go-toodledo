//go:generate stringer -type=TaskPriority

package taskpriorities

// TaskPriority ...
type TaskPriority int

// Negative ...
const (
	Negative TaskPriority = -1
	Low      TaskPriority = 0
	Medium   TaskPriority = 1
	High     TaskPriority = 2
	Top      TaskPriority = 3
)

// All ...
var All = []TaskPriority{
	Negative,
	Low,
	Medium,
	High,
	Top,
}

// Value2Type ...
func Value2Type(input int64) TaskPriority {
	for _, x := range All {
		if x == TaskPriority(input) {
			return x
		}
	}
	return Medium
}
