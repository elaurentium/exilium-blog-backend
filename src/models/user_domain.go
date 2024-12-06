package models

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/elaurentium/exilium-blog-backend/src/configuration/Errors"
)

// Domain cannot be exported !!
type UserDomain struct {
	Name string
	Email string
	Password string
	Age int8
}

// Create interface
type userDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}

func newUserDomain(email, password, name string, age int8) userDomainInterface {
	return &UserDomain{
		email,
		password,
		name,
		age,
	}
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}