package controllers

import (
	"net/http"

	validation "github.com/elaurentium/exilium-blog-backend/src/configuration/Validation"
	"github.com/elaurentium/exilium-blog-backend/src/models/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GET
func GetUserById(c *gin.Context) {}

func GetUserByEmail(c *gin.Context) {}

// POST
func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidadeUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := models.newUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created succefully",
		zap.String("Journey", "CreateUser"))

	c.String(http.StatusOK)
}

// DELETE
func DeleteUser(c *gin.Context) {}

// PUT
func UpdateUser(c *gin.Context) {}
