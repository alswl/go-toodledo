package main

import (
	"context"
	"fmt"
	"github.com/alswl/go-toodledo/toodledo"
	"log"
	"os"
)

func main() {
	accessToken := os.Getenv("TOODLEDO_ACCESS_TOKEN")

	if accessToken == "" {
		log.Fatal("Unauthorized: No TOODLEDO_ACCESS_TOKEN present")
	}

	ctx := context.Background()
	client := toodledo.NewClient(accessToken)
	folder, _, err := client.FolderService.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully get: %v\n", folder)
}
