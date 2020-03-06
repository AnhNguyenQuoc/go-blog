package controllers

import (
	"encoding/json"
	"github.com/AnhNguyenQuoc/go-blog/lib"
	"github.com/AnhNguyenQuoc/go-blog/models"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var todoService models.TodoService

type TodoResponse struct {
	Email string
	Todos []models.Todo
}

func TodoRouter(r *httprouter.Router, db *gorm.DB) {
	todoService = models.TodoService{DB: db}
	r.GET("/todos", TodosHandler)
	r.POST("/todo", CheckAuthenticate(TodoHandler))
	r.DELETE("/todo/:id", CheckAuthenticate(TodoHandler))
	r.PATCH("/todo/:id/:type", CheckAuthenticate(TodoHandler))
}

func TodosHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	user, _ := CurrentUser(w, r)
	response := TodoResponse{
		Email: user.Email,
		Todos: todoService.All(),
	}

	lib.ParseTemplate(w, "todo/index", response)
}

func TodoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == "POST" {
		var todo models.Todo
		json.NewDecoder(r.Body).Decode(&todo)
		todoService.Create(&todo)

		t, _ := json.Marshal(todo)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(t)
	} else if r.Method == "DELETE" {
		err := todoService.Delete(ps.ByName("id"))
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(http.StatusNoContent)
	} else if r.Method == "PATCH" {
		todoService.Update(ps.ByName("id"), ps.ByName("type"))

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ps.ByName("type")))
	}

}
