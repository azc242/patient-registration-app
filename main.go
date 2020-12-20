package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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
	_ = json.NewDecoder(r.Body).Decode(&patient) // decode json request body into Patient struct
	patient.ID = guuid.New().String()            // creates random ID for new patient
	patient.Time = time.Now().UTC().Format("2006-01-02 03:04:05")
	patients = append(patients, patient)

	// start db connection
	// Fetch environment MySQL environment variables from .env file
	username := goDotEnvVariable("MYSQL_USERNAME")
	password := goDotEnvVariable("MYSQL_PASSWORD")
	db, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/patientappdb")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// insert new patient into DB
	query := fmt.Sprintf("INSERT INTO patients VALUES ('%s','%s','%s','%s','%s','%s')", patient.ID, patient.Name, patient.DOB, patient.Phone, patient.Email, patient.Time)
	insert, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	// encode inserted patient as json and send back
	json.NewEncoder(w).Encode(patient)

}

// Validate admin user
func validateAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// params := mux.Vars(r)
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.Username == "0" && user.Password == "0Admin" {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	return
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func setDB() {
	// Fetch environment MySQL environment variables from .env file
	username := goDotEnvVariable("MYSQL_USERNAME")
	password := goDotEnvVariable("MYSQL_PASSWORD")

	// start MySQL connection
	db, err := sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/patientappdb")
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

	fmt.Print("Successfully inserted into user tables\n")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/patients", getPatients).Methods("GET")
	r.HandleFunc("/api/patients", createPatient).Methods("POST")
	r.HandleFunc("/api/login", validateAdmin).Methods("POST")

	// sets the db up with default admin
	setDB()

	// set up server on port 8000
	log.Fatal(http.ListenAndServe(":8080", r))
}
