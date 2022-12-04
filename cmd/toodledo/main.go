package main

import (
	"os"

	"github.com/alswl/go-toodledo/cmd/toodledo/root"
	"github.com/alswl/go-toodledo/pkg/common/logging"
)

func main() {
	err := logging.InitFactory("/tmp/toodledo", false, false)
	log := logging.ProvideLogger()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	root.Execute()
}
