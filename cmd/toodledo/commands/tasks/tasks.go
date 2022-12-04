package tasks

import (
	"github.com/alswl/go-toodledo/pkg/cmdutil"
	"github.com/spf13/cobra"
)

func NewCmd(f *cmdutil.Factory) *cobra.Command {
	list := NewListCmd(f)
	cmd := &cobra.Command{
		Use:   "task",
		Args:  list.Args,
		Short: "Manage toodledo tasks",
		Run:   list.Run,
	}
	cmd.AddCommand(
		list,
		NewCreateCmd(f),
		NewDeleteCmd(f),
		NewViewCmd(f),
		NewDoneCmd(f),
		NewReopenCmd(f),
		NewEditorCmd(f),
		NewEditCmd(f),
		NewStartCmd(f),
		NewStopCmd(f),
	)
	return cmd
}
