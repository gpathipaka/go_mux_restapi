package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server starting....")

	//Initialize Mux Router.
	r := mux.NewRouter()
	r.HandleFunc("/employees", createEmployee).Methods("POST")
	r.HandleFunc("/employees/{empId}", getEmployee).Methods("GET")
	r.HandleFunc("/employees", getAllEmployees).Methods("GET")
	r.HandleFunc("/employees", updateEmployee).Methods("PUT")
	r.HandleFunc("/employees", deleteEmployee).Methods("DELETE")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed to start ....", err)
	}

	log.Println("Server going down....")
}
