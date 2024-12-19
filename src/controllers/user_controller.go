package controllers

import (
	"github.com/elaurentium/exilium-blog-backend/src/models/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewUserControllerInterface(database *mongo.Database) UserControllerInterface {
	return &userControllerInterface{
		database: database,
	}
}

type UserControllerInterface interface {
	GetUserById(c *gin.Context)
	GetUserByEmail(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userControllerInterface struct {
	service  service.UserDomainService
	database *mongo.Database
}
