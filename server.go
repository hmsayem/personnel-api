package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/home/sayem/Downloads/firebase/employee-server.json")
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/employees", getEmployees).Methods("GET")
	router.HandleFunc("/employees", addEmployee).Methods("POST")
	log.Println("Server listening on port", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		return
	}

}
