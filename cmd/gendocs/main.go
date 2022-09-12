package main

import (
	"github.com/alswl/go-toodledo/cmd/toodledo/root"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra/doc"
)

func main() {
	err := doc.GenMarkdownTree(root.NewRootCmd(cmdutil.NewFactory()), "./docs/commands")
	if err != nil {
		logrus.WithError(err).Fatal("generate markdown failed")
	}
}
