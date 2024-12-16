package routers

import (

	"github.com/elaurentium/exilium-blog-backend/src/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(router*gin.RouterGroup, userController controllers.UserControllerInterface) {
	router.GET("/getUserById/:userId", userController.GetUserByEmail)
	router.GET("/getUserByEmail/:userEmail", userController.GetUserByEmail)
	router.POST("/createUser/", userController.CreateUser)
	router.PUT("/updateUser/:userId", userController.UpdateUser)
	router.DELETE("/deleteUser/:userId", userController.DeleteUser)
}