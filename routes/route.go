package routes

import (
	"github.com/AnhNguyenQuoc/go-blog/controllers"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
)

func InitRoute(r *httprouter.Router, db *gorm.DB) {
	controllers.UserRouter(r, db)
	controllers.AuthRouter(r, db)
	controllers.LayoutRouter(r, db)
}
