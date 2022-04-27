package main

import (
	"fmt"

	"github.com/marksman-jobs/backend/config"
	"github.com/marksman-jobs/backend/controller"
)

func main() {
	fmt.Println("Welcome to Marksman.com - Backend")
	fmt.Println("Starting processes...")

	conf := config.InitializeConfig()
	controller.InitializeServer(conf)
}
