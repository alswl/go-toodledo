//go:generate stringer -type=Priority

package tasks

import "strings"

// Priority ...
type Priority int

const (
	PriorityNegative Priority = -1
	// FIXME low is 0, 0 is zero value
	PriorityLow    Priority = 0
	PriorityMedium Priority = 1
	PriorityHigh   Priority = 2
	PriorityTop    Priority = 3
)

// PriorityAll ...
var PriorityAll = []Priority{
	PriorityNegative,
	PriorityLow,
	PriorityMedium,
	PriorityHigh,
	PriorityTop,
}

var PriorityMap = map[string]Priority{
	"negative": PriorityNegative,
	"low":      PriorityLow,
	"medium":   PriorityMedium,
	"high":     PriorityHigh,
	"top":      PriorityTop,
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

func PriorityString2Type(input string) Priority {
	for k, v := range PriorityMap {
		if k == strings.ToLower(input) {
			return v
		}
	}
	return PriorityMedium
}
