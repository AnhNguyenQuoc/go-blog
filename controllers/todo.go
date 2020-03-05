package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/AnhNguyenQuoc/go-blog/lib"
	"github.com/AnhNguyenQuoc/go-blog/models"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var todoService models.TodoService

type TodoResponse struct {
	Email string
	todos []models.Todo
}

func TodoRouter(r *httprouter.Router, db *gorm.DB) {
	todoService = models.TodoService{DB: db}
	r.GET("/todos", TodosHandler)
	r.POST("/todo", TodoHandler)
}

func TodosHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user, _ := CurrentUser(w, r, todoService.DB)
	response := TodoResponse{
		Email: user.Email,
		todos: todoService.All(),
	}

	lib.ParseTemplate(w, "todo/index", response)
}

func TodoHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	fmt.Println(todo)
	todoService.Create(&todo)

	fmt.Println("Created successfully")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
