package contexts

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/alswl/go-toodledo/cmd/toodledo/injector"
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(f *cmdutil.Factory) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Args:  cobra.ExactArgs(1),
		Short: "Delete a context",
		Example: heredoc.Doc(`
			$ toodledo context delete Work
		`),
		Run: func(cmd *cobra.Command, args []string) {
			_, err := injector.InitApp()
			if err != nil {
				logrus.Fatal("login required, using `toodledo auth login` to login.")
				return
			}
			// TODO app.svc
			svc, err := injector.InitContextService()
			if err != nil {
				logrus.Fatal(err)
				return
			}
			name := args[0]

			err = svc.Delete(name)
			if err != nil {
				logrus.Fatal(err)
				return
			}
			fmt.Println("done")
		},
	}
}
