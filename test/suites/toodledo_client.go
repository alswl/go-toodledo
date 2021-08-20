//+build integration

package suites

import (
	"github.com/alswl/go-toodledo/pkg/toodledo"
	"os"
)

func ClientForTest() *toodledo.Client {
	accessToken := os.Getenv("TOODLEDO_ACCESS_TOKEN")
	if accessToken == "" {
		panic("nil TOODLEDO_ACCESS_TOKEN")
	}

	client := toodledo.NewClient(accessToken)
	return client
}
