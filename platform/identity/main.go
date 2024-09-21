package main

import (
	"log"
	"pulselog/identity/config"
	"pulselog/identity/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnvironmentVars()

	db, err := config.InitDatabase()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	corsConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig))

	routes.SetupAuthRoutes(router, db)
	routes.SetupUserRoutes(router, db)
	routes.SetupProjectRoutes(router, db)
	routes.SetupProjectMemberRoutes(router, db)
	routes.SetupAPIKeysRoutes(router, db)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
