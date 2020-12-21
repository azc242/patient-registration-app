package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	guuid "github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// "time"

// Patient struct
type Patient struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	DOB     string `json:"dob"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Time    string `json:"time"`
}

// User struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Get all patients
func getPatients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

	patients, err := db.Query("SELECT* FROM patientappdb.patients")

	if err != nil {
		panic(err.Error())
	}

	// init patients as a slice Patient struct
	var patientsSlice []Patient
	for patients.Next() {
		var patient Patient

		err := patients.Scan(&patient.ID, &patient.Name, &patient.DOB, &patient.Phone, &patient.Email, &patient.Address, &patient.Time)
		if err != nil {
			panic(err.Error())
		}
		patientsSlice = append(patientsSlice, patient)
	}

	// encode slice containing Patient structs into json
	json.NewEncoder(w).Encode(patientsSlice)
}

// Create a patient
func createPatient(w http.ResponseWriter, r *http.Request) {
	// enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	var patient Patient
	_ = json.NewDecoder(r.Body).Decode(&patient) // decode json request body into Patient struct
	patient.ID = guuid.New().String()            // creates random ID for new patient
	patient.Time = time.Now().UTC().Format("2006-01-02 03:04:05")

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
	query := fmt.Sprintf("INSERT INTO patients VALUES ('%s','%s','%s','%s','%s','%s','%s')", patient.ID, patient.Name, patient.DOB, patient.Phone, patient.Email, patient.Address, patient.Time)
	insert, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	w.WriteHeader(http.StatusOK) // Sets 200 status
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

// Test POST requests
func testPOST(w http.ResponseWriter, r *http.Request) {
	// enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	var patient Patient
	_ = json.NewDecoder(r.Body).Decode(&patient) // decode json request body into Patient struct
	patient.ID = guuid.New().String()            // creates random ID for new patient
	patient.Time = time.Now().UTC().Format("2006-01-02 03:04:05")

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
	query := fmt.Sprintf("INSERT INTO patients VALUES ('%s','%s','%s','%s','%s','%s','%s')", patient.ID, patient.Name, patient.DOB, patient.Phone, patient.Email, patient.Address, patient.Time)
	insert, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	w.WriteHeader(http.StatusOK)

	// encode inserted patient as json and send back
	json.NewEncoder(w).Encode(patient)

}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/patients", getPatients).Methods("GET")
	r.HandleFunc("/api/patients", createPatient).Methods("POST")
	r.HandleFunc("/api/patients", createPatient).Methods("OPTIONS") // Helps with CORS
	r.HandleFunc("/api/login", validateAdmin).Methods("POST")
	// r.HandleFunc("/test", testPOST).Methods("POST")
	// r.HandleFunc("/test", testPOST).Methods("OPTIONS")

	// sets the db up with default admin
	// setDB()

	// set up server on port 8000
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"http://localhost:3000"}))(r)))
}

// Sets up the DB with default admin
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

	// Artificallu adds 2 test users to DB, good for testing GET on /api/patients endpoint
	var patient Patient
	patient.Name = "John Doe"
	patient.DOB = "12/1/1998"
	patient.Phone = "1234567890"
	patient.Email = "test123@fakemail.net"
	patient.Address = "99999 Carpenter St, Raleigh, NC 30021"
	patient.ID = guuid.New().String() // creates random ID for new patient
	patient.Time = time.Now().UTC().Format("2006-01-02 03:04:05")
	query := fmt.Sprintf("INSERT INTO patients VALUES ('%s','%s','%s','%s','%s','%s','%s')", patient.ID, patient.Name, patient.DOB, patient.Phone, patient.Email, patient.Address, patient.Time)
	insert, err := db.Query(query)

	patient.Name = "Bailey Thomas"
	patient.DOB = "3/9/1972"
	patient.Phone = "690-408-4188"
	patient.Email = "bt1972@gmail.com"
	patient.Address = "432 Garden Pkwy, Jacksonville, FL 21023"
	patient.ID = guuid.New().String() // creates random ID for new patient
	patient.Time = time.Now().UTC().Format("2006-01-02 03:04:05")
	query = fmt.Sprintf("INSERT INTO patients VALUES ('%s','%s','%s','%s','%s','%s','%s')", patient.ID, patient.Name, patient.DOB, patient.Phone, patient.Email, patient.Address, patient.Time)
	insert, err = db.Query(query)
	defer insert.Close()

	// Artificially set up the user admin
	// This user is the only use who can log in and see patients information
	// insert, err = db.Query("INSERT INTO users VALUES ('0', '0Admin')")

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Print("Successfully inserted into user tables\n")
}
