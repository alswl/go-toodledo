package auth

import "github.com/spf13/cobra"

var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage authentication",
}

func init() {
	AuthCmd.AddCommand(loginCmd, tokenCmd)
}
