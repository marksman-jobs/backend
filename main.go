package main

import (
	"fmt"

	"github.com/marksman-jobs/backend/config"
	"github.com/marksman-jobs/backend/controller"
	"github.com/marksman-jobs/backend/validation"
)

func main() {

	fmt.Println("Welcome to Marksman.com - Backend")
	fmt.Println("Starting processes...")

	conf := config.InitializeConfig()
	if validation.ValidateConfig(conf) {
		panic(fmt.Errorf("missing field(s) in config"))
	}

	db, err := config.InitializeDatabase(*conf)
	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}

	controller.InitializeServer(conf, db)

	// TODO: Add graceful shutdown here

}
