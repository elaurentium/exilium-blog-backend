package controllers

import (
	"fmt"

	rest_err "github.com/elaurentium/exilium-blog-backend/src/configuration/Errors"
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
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect filds, error=%s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}

// DELETE
func DeleteUser(c*gin.Context) {}

// PUT
func UpdateUser(c*gin.Context) {}