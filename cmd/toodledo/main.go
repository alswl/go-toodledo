package main

import (
	"fmt"
	"github.com/alswl/go-toodledo/cmd/toodledo/commands"
	"github.com/alswl/go-toodledo/pkg/common/logging"
	"os"
)

func main() {
	err := logging.InitFactory("/tmp/toodledo", false, false)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	commands.Execute()
}
