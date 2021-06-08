package main

import (
	"fmt"
	"log"

	"github.com/devops-kung-fu/domi/routes"
)

var err error

func main() {
	fmt.Println("domi - Policy-as-Code Enforcer")
	fmt.Println("Starting server on port 8080")

	r := routes.SetupRouter()
	err := r.Run()
	if err != nil {
		log.Fatal("An error has ocurred while starting the server.")
	}
}
