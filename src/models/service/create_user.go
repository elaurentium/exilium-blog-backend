package service

import (
	rest_err "github.com/elaurentium/exilium-blog-backend/src/configuration/Errors"
	"github.com/elaurentium/exilium-blog-backend/src/models"
	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

func (ud *userDomainService) CreateUser(userDomain models.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init create user model", zap.String("Jorney", "CreateUser"))
	userDomain.EncryptPassword()
	return nil
}

func (ud *userDomainService) UpdateUser(id string) *rest_err.RestErr {
	logger.Info("Init update user model", zap.String("Jorney", "UpdateUser"))
	return nil
}

func (ud *userDomainService) FindUser(id string, userDomain models.UserDomainInterface) (*userDomainService, *rest_err.RestErr) {
	logger.Info("Init find user model", zap.String("Jorney", "FindUser"))
	return nil, nil
}

func (ud *userDomainService) DeleteUser(id string) *rest_err.RestErr {
	logger.Info("Init delete user model", zap.String("Jorney", "DeleteUser"))
	return nil
}
