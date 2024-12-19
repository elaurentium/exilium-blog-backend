package controllers

import (
	"net/http"

	validation "github.com/elaurentium/exilium-blog-backend/src/configuration/Validation"
	"github.com/elaurentium/exilium-blog-backend/src/models"
	"github.com/elaurentium/exilium-blog-backend/src/models/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

// GET
func (uc *userControllerInterface) GetUserById(c *gin.Context) {}

func (uc *userControllerInterface) GetUserByEmail(c *gin.Context) {}

// POST
func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidadeUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := models.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created succefully",
		zap.String("Journey", "CreateUser"))

	c.JSON(http.StatusOK, "")
}

// DELETE
func (uc *userControllerInterface) DeleteUser(c *gin.Context) {}

// PUT
func (uc *userControllerInterface) UpdateUser(c *gin.Context) {}
