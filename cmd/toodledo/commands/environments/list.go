package environments

import (
	"fmt"

	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/alswl/go-toodledo/pkg/render"
	"github.com/alswl/go-toodledo/pkg/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewListCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Args:    cobra.NoArgs,
		Aliases: []string{"ls"},
		Short:   "List environments",
		Example: heredoc.Doc(`
			$ toodledo environment list
`),
		Run: func(cmd *cobra.Command, args []string) {
			cSrv := services.NewEnvironmentService()
			cks, err := cSrv.ListAll()
			if err != nil {
				logrus.Error(err)
				return
			}

			_, _ = fmt.Fprintln(f.IOStreams.Out, render.Environments(cks))
		},
	}
}
