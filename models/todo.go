package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Todo struct {
	gorm.Model
	Name string `gorm:"not_null" json:"name"`
	Done bool   `gorm:"default:false" json:"done"`
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
	t.DB.Table("todos").Order("created_at asc").Find(&todos)

	return todos
}

func (t TodoService) Delete(id string) error {
	idTodo, _ := strconv.Atoi(id)
	if result := t.DB.Table("todos").Where("id = ?", idTodo).Delete(&Todo{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (t TodoService) Update(id, done string) error {
	idTodo, _ := strconv.Atoi(id)
	status, _ := strconv.ParseBool(done)
	if result := t.DB.Table("todos").Where("id = ?", idTodo).Update("done", status)
		result.Error != nil {
		return result.Error
	}

	return nil
}
