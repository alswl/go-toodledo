package main

import (
	"fmt"
	"os"
	"net/http"
	"./toodledo"
)

func main() {
	fmt.Println("Hello, GO !")
	app_id := os.Getenv("TOODLE_APP_ID")
	fmt.Println(app_id)

	client := http.Client{}
	toodledo.NewClient(&client)
}
