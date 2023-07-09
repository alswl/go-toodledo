package utils

import (
	"strconv"
	"time"
)

// DefaultTimeZone is the default timezone
// FIXME using configuration time zone.
var DefaultTimeZone = asiaShanghaiTimeZone
var asiaShanghaiTimeZone = time.FixedZone("CST", 8*3600)

func Bool2int(input bool) int {
	var output int
	if input {
		output = 1
	} else {
		output = 0
	}
	return output
}

func Bool2ints(input bool) (output string) {
	return strconv.Itoa(Bool2int(input))
}
