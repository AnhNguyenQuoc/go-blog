package main

import (
	"fmt"
	"github.com/AnhNguyenQuoc/go-blog/routes"
	"log"
	"net/http"
	"os"

	"github.com/AnhNguyenQuoc/go-blog/lib"
	"github.com/AnhNguyenQuoc/go-blog/migrate"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
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
	// Initialize router
	routes.InitRoute(router, db)
	// Static file router
	router.ServeFiles("/static/*filepath", http.Dir("assets/"))
	// Start server
	fmt.Println("Listening from server...")
	log.Fatal(http.ListenAndServe(lib.GetPort(), router))
}
