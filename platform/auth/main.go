package main

import (
	"log"
	"pulselog/auth/config"
	"pulselog/auth/routes/v1"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.InitDatabase()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	routes.SetupV1Routes(router, db)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
