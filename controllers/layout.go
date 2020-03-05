package controllers

import (
	"github.com/AnhNguyenQuoc/go-blog/lib"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var layoutService LayoutService

type LayoutService struct {
	DB *gorm.DB
}

func LayoutRouter(r *httprouter.Router, db *gorm.DB) {
	layoutService = LayoutService{DB: db}
	r.GET("/", HelloWorld)
}

func HelloWorld(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	user, ok := CurrentUser(w, r, layoutService.DB)
	if !ok {
		lib.ParseTemplate(w, "layout/index", nil)
		return
	}

	lib.ParseTemplate(w, "layout/index", user)
}
