package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// "time"

// Patient struct
type Patient struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	DOB   string `json:"dob"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Time  string `json:"time"`
}

// User struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

// Validate admin user
func validateAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.Username == "0" && user.Password == "Admin0" {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	return
}

func main() {
	r := mux.NewRouter()

	patients = append(patients, Patient{ID: "1", Name: "John Doe", DOB: "12/1/1998", Phone: "1234567890", Email: "test@123.net", Time: time.Now().String()})
	patients = append(patients, Patient{ID: "2", Name: "David Peterson", DOB: "6/2/1980", Phone: "9876543210", Email: "dp@gmail.com", Time: time.Now().String()})

	r.HandleFunc("/api/patients", getPatients).Methods("GET")
	r.HandleFunc("/api/patients", createPatient).Methods("POST")
	r.HandleFunc("/api/login", validateAdmin).Methods("POST")

	// start MySQL connection
	db, err := sql.Open("mysql", "root:Clutch4405!@tcp(127.0.0.1:3306)/patientappdb")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Artificially set up the user admin
	// this user is the only use who can log in and see patients information
	insert, err := db.Query("INSERT INTO users VALUES ('0', '0Admin')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Print("Successfully inserted into user tables")

	// set up server on port 8000
	log.Fatal(http.ListenAndServe(":8080", r))
}

// mysql 0Admin
