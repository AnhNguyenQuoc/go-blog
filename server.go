package main

import (
	"fmt"
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

func init() {
	godotenv.Load()
	dbConfig = migrate.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	_, err := migrate.InitDB(dbConfig)
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	os.Setenv("PORT", "3000") //TODO: remove when push to heroku

	router := httprouter.New()

	router.GET("/", HelloWorld)
	// Static file router
	router.ServeFiles("/static/*filepath", http.Dir("assets/"))

	fmt.Println("Listening from server...")
	log.Fatal(http.ListenAndServe(lib.GetPort(), router))
}

func HelloWorld(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	lib.ParseTemplate(w, "views/layout/index", map[string]string{"Name": "Anh"})
}
