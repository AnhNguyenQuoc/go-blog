package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"regexp"
)

var regexEmail = regexp.MustCompile(`^\w.+@(\w)+\.\w+$`)

type User struct {
	gorm.Model
	Username string `gorm:"not_null"`
	Password string `gorm:"not_null"`
	Email    string `gorm:"type:varchar(100);unique_index;unique;not_null"`
}

func (user *User) CreateUser(db *gorm.DB) error {
	dbc := db.Create(&User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
	})
	if dbc.Error != nil {
		return errors.New("data need to unique. Please check again")
	}
	return nil
}

func (user User) Validate() map[string]string {
	err := map[string]string{}

	if user.Username == "" {
		err["name"] = "The username field is required"
	}

	if user.Email == "" {
		err["email"] = "The email field is required"
	}

	if !regexEmail.MatchString(user.Email) {
		err["emailFormat"] = "The email field is not format email"
	}

	if user.Password == "" {
		err["password"] = "The password field is required"
	}

	return err
}
