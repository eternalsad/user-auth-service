package main

import (
	"log"
	"os"

	"github.com/eternalsad/user-auth-service/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
