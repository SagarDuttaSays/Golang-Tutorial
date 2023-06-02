package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4003", r))
}
func greeter() {
	fmt.Println("Hello... seasonal greetings!")
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>I love my INDIA</h1>"))
}
