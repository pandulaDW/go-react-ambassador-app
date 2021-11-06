package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id           uint
	FirstName    string
	LastName     string
	Email        string
	Password     []byte
	IsAmbassador bool
}

func (user *User) SetPassword(password []byte) []byte {
	encryptedPassword, _ := bcrypt.GenerateFromPassword(password, 12)
	return encryptedPassword
}

func (user *User) ComparePassword(password []byte) error {
	return bcrypt.CompareHashAndPassword(user.Password, password)
}
