//go:build integration
// +build integration

package suites

import (
	"github.com/alswl/go-toodledo/pkg/auth"
	"github.com/alswl/go-toodledo/pkg/toodledo"
	"github.com/go-openapi/runtime"
	"github.com/spf13/viper"
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

func AuthForTest() runtime.ClientAuthInfoWriter {
	home, _ := os.UserHomeDir()
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".toodledo")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	a, err := auth.ProvideSimpleAuth()
	if err != nil {
		panic(err)
	}
	return a
}
