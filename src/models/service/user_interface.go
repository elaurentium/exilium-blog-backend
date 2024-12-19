package service

import (
	"github.com/HunCoding/meu-primeiro-crud-go/src/model/repository"
	rest_err "github.com/elaurentium/exilium-blog-backend/src/configuration/Errors"
	"github.com/elaurentium/exilium-blog-backend/src/models"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

// Create interface
type UserDomainService interface {
	CreateUser(models.UserDomainInterface) *rest_err.RestErr
	UpdateUser(string, models.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*models.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
