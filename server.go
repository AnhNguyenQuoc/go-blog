package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("PORT", "3000") //TODO: remove when push to heroku

	router := httprouter.New()

	router.GET("/", HelloWorld)
	// Static file router
	router.ServeFiles("/static/*filepath", http.Dir("assets/"))

	fmt.Println("Listening from server...")
	log.Fatal(http.ListenAndServe(GetPort(), router))
}

func HelloWorld(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello 33333!")
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Err when open PORT env: ", port)
	}

	return ":" + port
}
