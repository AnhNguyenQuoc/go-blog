package models

import (
	"errors"
	"github.com/AnhNguyenQuoc/go-blog/lib"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not_null"`
	Password string `gorm:"not_null"`
	Email    string `gorm:"type:varchar(100);unique_index;unique;not_null"`
}

type UserService struct {
	DB *gorm.DB
}

func (r *UserService) CreateUser(user *User) error {
	dbc := r.DB.Create(user)
	if dbc.Error != nil {
		return errors.New("data need to unique. Please check again")
	}
	return nil
}

func (u User) Validate() map[string]string {
	err := map[string]string{}

	if u.Username == "" {
		err["name"] = "The username field is required"
	}

	if u.Email == "" {
		err["email"] = "The email field is required"
	}

	if !lib.RegexEmail.MatchString(u.Email) {
		err["emailFormat"] = "The email field is not format email"
	}

	if u.Password == "" {
		err["password"] = "The password field is required"
	}

	return err
}

func (u *User) BeforeSave() (err error) {
	u.Password, err = lib.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
