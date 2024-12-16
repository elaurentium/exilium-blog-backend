package service

import (
	rest_err "github.com/elaurentium/exilium-blog-backend/src/configuration/Errors"
	"github.com/elaurentium/exilium-blog-backend/src/models"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
	
}

// Create interface
type UserDomainService interface {
	CreateUser(models.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, models.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*models.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}