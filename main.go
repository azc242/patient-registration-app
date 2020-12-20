package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
)

// "time"

// Patient struct
type Patient struct {
	ID    string `json:"id"`
	Name  string `json:"namne"`
	DOB   string `json:"dob"`
	Phone int    `json:"phone"`
	Email string `json:"email"`
	Time  string `json:"time"`
}

// init patients as a slice Patient struct
var patients []Patient

// Get all patients
func getPatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// encode slice containing patients into json
	json.NewEncoder(w).Encode(patients)
}

// Create a patient
func createPatient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var patient Patient
	_ = json.NewDecoder(r.Body).Decode(&patient) // decode json requesy body into Patient struct
	patient.ID = guuid.New().String()            // creates random ID for new patient
	patients = append(patients, patient)
	json.NewEncoder(w).Encode(patient)
}

func main() {
	r := mux.NewRouter()

	patients = append(patients, Patient{ID: "1", Name: "John Doe", DOB: "12/1/1998", Phone: 1234567890, Email: "test@123.net", Time: time.Now().String()})
	patients = append(patients, Patient{ID: "2", Name: "David Peterson", DOB: "6/2/1980", Phone: 9876543210, Email: "dp@gmail.com", Time: time.Now().String()})

	r.HandleFunc("/api/patients", getPatients).Methods("GET")
	r.HandleFunc("/api/paients", createPatient).Methods("POST")

	// set up server on port 8000
	log.Fatal(http.ListenAndServe(":8080", r))
}
