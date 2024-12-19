package main

import (
	"context"
	"log"

	"github.com/elaurentium/exilium-blog-backend/src/configuration/Logs"
	"github.com/elaurentium/exilium-blog-backend/src/configuration/database/mongodb"
	"github.com/elaurentium/exilium-blog-backend/src/controllers"
	"github.com/elaurentium/exilium-blog-backend/src/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	Logs.Info("About to start application")
	godotenv.Load()

	database, err := mongodb.NewMongoDBConnection(context.Background())

	if err != nil {
		log.Fatal(
			"Error trying to connect to database",
			err.Error(),
		)
		return
	}

	// init dependencies
	userControler := controllers.NewUserControllerInterface(database)

	router := gin.Default() // boot router with logger and recovery middlewares

	routers.InitRouter(&router.RouterGroup, userControler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
