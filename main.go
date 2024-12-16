package main

import (
	"log"

	"github.com/elaurentium/exilium-blog-backend/src/configuration/Logs"
	"github.com/elaurentium/exilium-blog-backend/src/configuration/database/mongodb"
	"github.com/elaurentium/exilium-blog-backend/src/controllers"
	"github.com/elaurentium/exilium-blog-backend/src/models/service"
	"github.com/elaurentium/exilium-blog-backend/src/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// init dependencies
	service :=  service.NewUserDomainService()
	userControler :=  controllers.NewUserControllerInterface(service)

	mongodb.InitConnection()

	Logs.Info("About to start application")
	router := gin.Default() // boot router with logger and recovery middlewares

	routers.InitRouter(&router.RouterGroup, userControler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}