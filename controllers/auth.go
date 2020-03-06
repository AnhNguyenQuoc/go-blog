package controllers

import (
	"github.com/AnhNguyenQuoc/go-blog/lib"
	"github.com/AnhNguyenQuoc/go-blog/models"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var Auth AuthService

type AuthRequest struct {
	Email    string
	Password string
}

type AuthSession struct {
	Login    bool
	UserName string
}

type AuthResponseError struct {
	Error []string
	Email string
}

type AuthService struct {
	DB *gorm.DB
}

func AuthRouter(r *httprouter.Router, db *gorm.DB) {
	Auth = AuthService{DB: db}
	r.GET("/login", LoginHandler)
	r.POST("/login", LoginHandler)
	r.POST("/logout", LogoutHandler)
}

func LoginHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		lib.ParseTemplate(w, "auth/login", nil)
	} else if r.Method == "POST" {
		var user models.User
		r.ParseForm()

		errors := AuthResponseError{Error: validateAuth(r)}
		if len(errors.Error) != 0 {
			w.WriteHeader(http.StatusBadRequest)
			lib.ParseTemplate(w, "/auth/login", errors)
			return
		}

		Auth.DB.Table("users").Where("email = ?", r.FormValue("email")).Find(&user)
		if !lib.AuthenticatePassword(r.FormValue("password"), user.Password) {
			w.WriteHeader(http.StatusBadRequest)
			errors.Error = append(errors.Error, "Authenticate fail. Please check again")
			errors.Email = ""
			lib.ParseTemplate(w, "/auth/login", errors)
			return
		}

		Login(&user, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := http.Cookie{
		Name:   "user_email",
		MaxAge: -1,
	}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func validateAuth(r *http.Request) []string {
	var result []string
	if r.FormValue("email") == "" {
		result = append(result, "The email field is required")
	}

	if r.FormValue("password") == "" {
		result = append(result, "The password field is required")
	}

	if !lib.RegexEmail.MatchString(r.FormValue("email")) {
		result = append(result, "The email field is wrong format")
	}

	return result
}

func CurrentUser(w http.ResponseWriter, r *http.Request) (models.User, bool) {
	userEmail, err := r.Cookie("user_email")
	var user models.User
	if err != nil {
		return user, false
	}

	if result := Auth.DB.Table("users").Where("email = ?", userEmail.Value).Find(&user); result.Error != nil {
		return user, false
	}

	return user, true
}

func CheckAuthenticate(f httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		_, ok := CurrentUser(w, r)
		if !ok {
			http.Redirect(w, r, "/", http.StatusUnauthorized)
			return
		}
		f(w, r, ps)
	}
}
