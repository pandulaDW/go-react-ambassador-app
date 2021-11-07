package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" gorm:"unique"`
	Password     []byte `json:"-"`
	IsAmbassador bool   `json:"-"`
}

func (user *User) SetPassword(password []byte) {
	encryptedPassword, _ := bcrypt.GenerateFromPassword(password, 12)
	user.Password = encryptedPassword
}

func (user *User) ComparePassword(password []byte) error {
	return bcrypt.CompareHashAndPassword(user.Password, password)
}
