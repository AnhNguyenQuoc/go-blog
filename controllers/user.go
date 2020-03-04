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

type UserController struct {
	db *gorm.DB
}

type ErrorMessage struct {
	Error map[string]string
}

func UserRouter(r *httprouter.Router, db *gorm.DB) {
	r.GET("/register", UserController{db: db}.Create)
	r.POST("/register", UserController{db: db}.Create)
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
			err := user.CreateUser(u.db)
			if err != nil {
				errorMessage := ErrorMessage{Error: map[string]string{"Error": fmt.Sprintln(err)}}
				err = lib.ParseTemplate(w, "/users/create", errorMessage)
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
}
