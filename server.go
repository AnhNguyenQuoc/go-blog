package main

import (
	"fmt"
	"github.com/AnhNguyenQuoc/go-blog/controllers"
	"github.com/AnhNguyenQuoc/go-blog/lib"
	"github.com/AnhNguyenQuoc/go-blog/migrate"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

var dbConfig migrate.DBConfig
var db *gorm.DB
var err error

func init() {
	godotenv.Load()
	dbConfig = migrate.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err = migrate.InitDB(dbConfig)
	if err != nil {
		log.Panic(err)
	}

}

func main() {
	defer db.Close()
	router := httprouter.New()

	os.Setenv("PORT", "3000") //TODO: remove when push to heroku

	// Setup custom message validate
	lib.CustomMessageErrorValidate()

	// Other router
	router.GET("/", HelloWorld)
	// User router
	controllers.UserRouter(router, db)
	// Static file router
	router.ServeFiles("/static/*filepath", http.Dir("assets/"))
	// Start server
	fmt.Println("Listening from server...")
	log.Fatal(http.ListenAndServe(lib.GetPort(), router))
}

func HelloWorld(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	err = lib.ParseTemplate(w, "layout/index", map[string]string{"Name": "Anh"})
	if err != nil {
		log.Fatal(err)
	}
}
