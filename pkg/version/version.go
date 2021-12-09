package version

import (
	"fmt"
	"runtime"
)

// Version ...
var (
	Version   = "0.0.0"
	Commit    = "UNKNOWN"
	Package   = "github.com/alswl/toodledo"
	BuildDate = "UNKNOWN"
	GoVersion = runtime.Version()
)

// Message ...
func Message() string {
	const format = `toodledo:   %s (Revision: %s)
package:    %s
build date: %s
go version: %s
`
	return fmt.Sprintf(format, Version, Commit, Package, BuildDate, GoVersion)
}
