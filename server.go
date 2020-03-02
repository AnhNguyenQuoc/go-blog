package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("PORT", "3000")

	http.HandleFunc("/", HelloWorld)
	fmt.Println("Listening from server...")

	log.Fatal(http.ListenAndServe(GetPort(), nil))
}

func HelloWorld(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello 22222!")
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Err when open PORT env: ", port)
	}

	return ":" + port
}
