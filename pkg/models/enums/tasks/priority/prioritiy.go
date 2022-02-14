//go:generate stringer -type=Priority
package priority

import (
	"strings"
)

// Priority ...
type Priority int

const (
	Negative Priority = -1
	// TODO low is 0, 0 is zero value
	Low    Priority = 0
	Medium Priority = 1
	High   Priority = 2
	Top    Priority = 3
)

// PriorityAll ...
var PriorityAll = []Priority{
	Negative,
	Low,
	Medium,
	High,
	Top,
}

var PriorityMap = map[string]Priority{
	"negative": Negative,
	"low":      Low,
	"medium":   Medium,
	"high":     High,
	"top":      Top,
}

// PriorityValue2Type ...
func PriorityValue2Type(input int64) Priority {
	for _, x := range PriorityAll {
		if x == Priority(input) {
			return x
		}
	}
	return Medium
}

func PriorityString2Type(input string) Priority {
	for k, v := range PriorityMap {
		if k == strings.ToLower(input) {
			return v
		}
	}
	return Medium
}
