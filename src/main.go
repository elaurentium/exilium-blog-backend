package main

import (
	"log"

	"github.com/elaurentium/exilium-blog-backend/src/configuration/Logs"
	"github.com/elaurentium/exilium-blog-backend/src/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	Logs.Info("About to start application")
	router := gin.Default() // boot router with logger and recovery middlewares

	routers.InitRouter(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}