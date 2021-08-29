package auth

import (
	"fmt"
	"github.com/alswl/go-toodledo/pkg/auth"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use: "login",
	Run: func(cmd *cobra.Command, args []string) {
		conf := auth.ProvideOAuth2Config()
		u2 := conf.AuthCodeURL("state")
		// TODO
		//pkg.OpenBrowser(u2)
		fmt.Printf("login in your browser in %s\n", u2)
		fmt.Println("login in your browser, then copy the url to clipboard and run `toodledo config callback YOUR-URL-AFTER-LOGIN`")
	},
}
