package pkg

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os/exec"
	"runtime"
	"strconv"
)

// Bool2int ...
func Bool2int(input bool) (output int) {
	if input {
		output = 1
	} else {
		output = 0
	}
	return
}

// Bool2ints ...
func Bool2ints(input bool) (output string) {
	return strconv.Itoa(Bool2int(input))
}

// OpenBrowser ...
func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		logrus.Error(err)
	}
}
