package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// "encoding/json"
// "log"
// "math/rand"
// "net/http"
// "strconv"
// "time"

type Patient struct {
	ID    string `json:"id"`
	Name  string `json:"namne"`
	DOB   string `json:"dob"`
	Phone int    `json:"phone"`
	Email string `json:"email"`
	Time  string `json:"time"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/patients", getPatients).Methods("GET")
	r.HandleFunc("/api/paients", createPaient).Methods("POST")

	// set up server on port 8000
	log.Fatal(http.ListenAndServe(":8080", r))
}
