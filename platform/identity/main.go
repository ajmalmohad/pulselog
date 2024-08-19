package main

import (
	"log"
	"pulselog/identity/config"
	"pulselog/identity/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnvironmentVars()

	db, err := config.InitDatabase()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	routes.SetupAuthRoutes(router, db)
	routes.SetupUserRoutes(router, db)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
