package models

import (
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

func (s *UserService) CreateUser(_ User) (User, error) {
	u := User{
		Username: "test",
		Password: "test",
		Email:    "test",
	}

	s.DB.Create(&u)

	return u, nil
}
