package models

import (
	rest_err "github.com/elaurentium/exilium-blog-backend/src/configuration/Errors"
	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()


func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init create user model", zap.String("Jorney", "CreateUser"))
	ud.EncryptPassword()
	return nil
}

func (ud *UserDomain) UpdateUser(id string) *rest_err.RestErr {
	logger.Info("Init update user model", zap.String("Jorney", "UpdateUser"))
	return nil
}

func (ud *UserDomain) FindUser(id string) (*UserDomain, *rest_err.RestErr) {
	logger.Info("Init find user model", zap.String("Jorney", "FindUser"))
	return nil, nil
}

func (ud *UserDomain) DeleteUser(id string) *rest_err.RestErr {
	logger.Info("Init delete user model", zap.String("Jorney", "DeleteUser"))
	return nil
}

