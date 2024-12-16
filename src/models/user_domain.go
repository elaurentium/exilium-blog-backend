package models

import (
	"crypto/md5"
	"encoding/hex"
)

// Domain cannot be exported !!
type userDomain struct {
	name string
	email string
	password string
	age int8
}

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	EncryptPassword()
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	//TODO
	return &userDomain{email, password, name, age}
}

func (ud* userDomain) GetEmail() string {
	return ud.email
}
func (ud* userDomain) GetPassword() string {
	return ud.password
}
func (ud* userDomain) GetName() string {
	return ud.name
}
func (ud* userDomain) GetAge() int8{
	return ud.age
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}