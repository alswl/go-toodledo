package auth

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "auth <command>",
	Short: "Manage authentication",
}

func init() {
	Cmd.AddCommand(loginCmd, tokenCmd)
}
