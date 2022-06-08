package tasks

import (
	"fmt"
	utilsos "github.com/alswl/go-toodledo/pkg/utils/os"
	"github.com/spf13/cobra"
	"strconv"
)

var browserCmd = &cobra.Command{
	Use:  "browser",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO support task name, first match
		id, _ := strconv.Atoi(args[0])
		taskURL := "https://www.toodledo.com/tasks/index.php?#task_%d"
		url := fmt.Sprintf(taskURL, id)
		_ = utilsos.OpenInBrowser(url)
	},
}

func init() {
	TaskCmd.AddCommand(browserCmd)
}
