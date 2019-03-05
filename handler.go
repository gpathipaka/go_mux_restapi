package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

//Employee is
type Employee struct {
	EmpID          string
	EmpName, Email string
	Sal            int
}

var employees []Employee

func createEmployee(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	log.Println(req.Body)
	_ = json.NewDecoder(req.Body).Decode(&emp)
	emp.EmpID = strconv.Itoa(rand.Intn(1000000))
	employees = append(employees, emp)
	log.Println("Saved the new Employee ", emp)
}

func getEmployee(w http.ResponseWriter, req *http.Request) {

}

func getAllEmployees(w http.ResponseWriter, req *http.Request) {
	log.Println("Returing all the employees....")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func updateEmployee(w http.ResponseWriter, req *http.Request) {

}

func deleteEmployee(w http.ResponseWriter, req *http.Request) {

}
