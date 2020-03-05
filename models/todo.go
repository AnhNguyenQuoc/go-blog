package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Name string `gorm:"not_null"`
	Done bool   `gorm:"default:false"`
}

type TodoService struct {
	DB *gorm.DB
}

func (t TodoService) Create(data *Todo) error {
	dbc := t.DB.Create(data)
	if dbc.Error != nil {
		return errors.New("data need to unique. Please check again")
	}
	return nil
}

func (t TodoService) All() []Todo {
	var todos []Todo
	t.DB.Table("todos").Find(&todos)

	return todos
}