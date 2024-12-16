package controllers

import (
	"github.com/elaurentium/exilium-blog-backend/src/models/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(service service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service,
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
	service service.UserDomainService
}