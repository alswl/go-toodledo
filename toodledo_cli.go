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

	client := toodledo.NewClient(accessToken)
	
	//testGet(client)
	//testAdd(client)
	//testEdit(client)
	testDelete(client)
}

func testGet(client *toodledo.Client) {
	ctx := context.Background()
	folders, _, err := client.FolderService.Get(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, folder := range folders {
		fmt.Printf("Successfully get: %v\n", folder)
	}
}

func testAdd(client *toodledo.Client) {
	ctx := context.Background()
	folder, resp, err := client.FolderService.Add(ctx, "test-abc")
	if err != nil {
		log.Fatal(err, resp)
		return
	}
	fmt.Printf("Successfully get: %v\n", folder)
}

func testEdit(client *toodledo.Client) {
	ctx := context.Background()
	folder, resp, err := client.FolderService.Edit(ctx, 9096513, "test-cba")
	if err != nil {
		log.Fatal(err, resp)
		return
	}
	fmt.Printf("Successfully get: %v\n", folder)
}

func testDelete(client *toodledo.Client) {
	ctx := context.Background()
	resp, err := client.FolderService.Delete(ctx, 9096513)
	if err != nil {
		log.Fatal(err, resp)
		return
	}
}
