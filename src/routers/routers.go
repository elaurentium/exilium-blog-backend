package routers

import (
	"github.com/elaurentium/exilium-blog-backend/src/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(router*gin.RouterGroup) {
	router.GET("/getUserById/:userId", controllers.GetUserByEmail)
	router.GET("/getUserByEmail/:userEmail", controllers.GetUserByEmail)
	router.POST("/createUser/", controllers.CreateUser)
	router.PUT("/updateUser/:userId", controllers.UpdateUser)
	router.DELETE("/deleteUser/:userId", controllers.DeleteUser)
}