package main

import (
	"github.com/alswl/go-toodledo/pkg/toodledo"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	accessToken := os.Getenv("TOODLEDO_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("Unauthorized: No TOODLEDO_ACCESS_TOKEN present")
	}

	toodledo.NewClient(accessToken)
}
