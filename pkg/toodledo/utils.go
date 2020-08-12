package toodledo

import "strconv"

func bool2int(input bool) (output int) {
	if input {
		output = 1
	} else {
		output = 0
	}
	return
}

func bool2ints(input bool) (output string) {
	return strconv.Itoa(bool2int(input))
}
