package controllers

import (
	"fmt"
	validation "github.com/elaurentium/exilium-blog-backend/src/configuration/Validation"
	"github.com/elaurentium/exilium-blog-backend/src/models/request"
	"github.com/gin-gonic/gin"
)

// GET
func GetUserById(c*gin.Context) {}

func GetUserByEmail(c*gin.Context) {}

// POST
func CreateUser(c*gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidadeUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}

// DELETE
func DeleteUser(c*gin.Context) {}

// PUT
func UpdateUser(c*gin.Context) {}