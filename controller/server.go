package controller

import (
	"fmt"

	"github.com/marksman-jobs/backend/config"
)

func InitializeServer(*config.Config) {
	fmt.Println("Server Initialized")
}
