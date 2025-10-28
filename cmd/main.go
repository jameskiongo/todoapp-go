package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/add", createTask).Methods("POST")
	r.HandleFunc("/task/{id}", getTasks).Methods("GET")
	r.HandleFunc("/task/{id}", updateTask).Methods("PUT")
	r.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")
	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", r)
	log.Fatal(err)
}
