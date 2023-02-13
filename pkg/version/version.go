package version

import (
	"fmt"
	"runtime"
)

var (
	// Version is the version of the build.
	// it will be overwritten automatically by the build system.
	Version   = "0.0.0"
	Commit    = "UNKNOWN"
	Package   = "github.com/alswl/toodledo"
	BuildDate = "UNKNOWN"
)

var GoVersion = runtime.Version()

func Message() string {
	const format = `toodledo:   %s (Revision: %s)
package:    %s
build date: %s
go version: %s
`
	return fmt.Sprintf(format, Version, Commit, Package, BuildDate, GoVersion)
}
