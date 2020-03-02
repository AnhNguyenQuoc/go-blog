package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("PORT", "3000")

	http.HandleFunc("/", helloWorld)
	fmt.Println("Listening from server...")

	log.Fatal(http.ListenAndServe(GetPort(), nil))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello there!")
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Err when open PORT env: ", port)
	}

	return ":" + port
}
