package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Setup the array of people.
var people []Person

// GetPeople returns a list of all people.
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// GetPerson returns a list of one person.
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

// CreatePerson adds a person to the list.
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePerson removes a person from the list.
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(people)
}

// Main method to serve up content and be able to hit routes.
func main() {
	router := mux.NewRouter()

	// Create functions that we're able to hit.
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	// Setup the fake data.
	people = append(people, Person{ID: "1", FirstName: "Mike", LastName: "A", Address: &Address{City: "Rochester", State: "New York"}})
	people = append(people, Person{ID: "2", FirstName: "Mike", LastName: "B", Address: &Address{City: "Rochester", State: "Minnesota"}})
	people = append(people, Person{ID: "3", FirstName: "Mike", LastName: "C", Address: &Address{City: "Rochester", State: "New Jersey"}})
	people = append(people, Person{ID: "4", FirstName: "Mike", LastName: "D"})

	// Setup the server!
	log.Fatal(http.ListenAndServe(":8000", router))
}
