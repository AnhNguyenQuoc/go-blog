package controllers

import (
	"github.com/AnhNguyenQuoc/go-blog/lib"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	err := lib.ParseTemplate(w, "layout/index", nil)
	if err != nil {
		log.Fatal(err)
	}
}
