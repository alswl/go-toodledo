package browse

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	utilsos "github.com/alswl/go-toodledo/pkg/utils/os"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strconv"
)

type Options struct {
	// TODO support title
	selector string

	task bool
	// TODO folder goal context
}

var opts = &Options{}

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "browse",
		Args:  cobra.MaximumNArgs(1),
		Short: "Browse toodledo in browser",
		Example: heredoc.Doc(`
toodledo browse
toodledo browse --task 288246017
`),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				opts.selector = args[0]
			}

			url := ""
			if opts.task {
				id, err := strconv.Atoi(opts.selector)
				if err != nil {
					logrus.Fatal("invalid task id")
					return
				}
				url = fmt.Sprintf("https://www.toodledo.com/tasks/index.php?#task_%d", id)
			} else {
				url = "https://www.toodledo.com/"
			}
			err := utilsos.OpenInBrowser(url)
			if err != nil {
				logrus.WithError(err).Fatal("open browser failed")
				return
			}
		},
	}

	cmd.Flags().BoolVarP(&opts.task, "task", "t", false, "open task in browser")
	return cmd
}
