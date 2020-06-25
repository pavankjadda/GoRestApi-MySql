package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/employees", getAllEmployees)
	myRouter.HandleFunc("/employee", createNewEmployee).Methods("POST")
	myRouter.HandleFunc("/employee/{id}", deleteEmployee).Methods("DELETE")
	myRouter.HandleFunc("/employee/{id}", returnSingleEmployee)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnSingleEmployee(writer http.ResponseWriter, request *http.Request) {

}

func deleteEmployee(writer http.ResponseWriter, request *http.Request) {

}

func createNewEmployee(writer http.ResponseWriter, request *http.Request) {

}

func getAllEmployees(writer http.ResponseWriter, request *http.Request) {
	json.NewEncoder(writer).Encode("")
}

func main() {
	handleRequests()
}
