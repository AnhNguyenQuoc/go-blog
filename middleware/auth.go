package middleware

import (
	"github.com/AnhNguyenQuoc/go-blog/controllers"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func CheckAuthenticate(f httprouter.Handle, db *gorm.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		_, ok := controllers.CurrentUser(w, r, db)
		if !ok {
			http.Redirect(w, r, "/", http.StatusUnauthorized)
			return
		}
		f(w, r, ps)
	}
}
