package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "james"
	password = "james"
	dbname   = "todo"
)

func main() {
	r := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Connecting sql
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, error := sql.Open("postgres", psqlInfo)
	if error != nil {
		panic(error)
	}
	defer db.Close()

	error = db.Ping()
	if error != nil {
		panic(error)
	} else {
		fmt.Println("Database connected")
	}
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
