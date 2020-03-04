package controllers

import (
	"fmt"
	"github.com/AnhNguyenQuoc/go-blog/lib"
	"github.com/AnhNguyenQuoc/go-blog/models"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var userService *models.UserService

type ErrorMessage struct {
	Error map[string]string
}

type UserController struct {
}

func UserRouter(r *httprouter.Router, db *gorm.DB) {
	userService = &models.UserService{DB: db}
	r.GET("/register", UserController{}.Create)
	r.POST("/register", UserController{}.Create)
}

func (u UserController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		err := lib.ParseTemplate(w, "/users/create", nil)
		if err != nil {
			log.Fatal(err)
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		user := &models.User{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
			Email:    r.FormValue("email"),
		}
		err := user.Validate()
		if len(err) != 0 {
			errorMessage := ErrorMessage{err}
			err := lib.ParseTemplate(w, "/users/create", errorMessage)
			if err != nil {
				log.Fatal(err)
			}
			return
		} else {
			user.Password, _ = lib.HashPassword(r.FormValue("password"))
			err := userService.CreateUser(user)
			if err != nil {
				errorMessage := ErrorMessage{Error: map[string]string{"Error": fmt.Sprintln(err)}}
				err = lib.ParseTemplate(w, "/users/create", errorMessage)
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
}
