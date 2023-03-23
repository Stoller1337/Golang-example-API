package internal

import (
	"awesomeProject2/internal/entity"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

func RunServer() {
	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/create", createPerson).Methods("POST")
	router.HandleFunc("/get/{id}", getPersonByID).Methods("GET")
	router.HandleFunc("/get", getPerson).Methods("GET")
	router.HandleFunc("/update/{id}", updatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", deletPersonByID).Methods("DELETE")
	log.Println(http.ListenAndServe(":8090", router))
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//key := vars["id"]
	var person entity.Person
	res := Connector.First(person, 1)
	log.Println(res.Error)
	if res.Error != nil {
		panic(res.Error)
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(person)
	if err != nil {
		log.Fatalln(err)
	}
}
func createPerson(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := io.ReadAll(r.Body)
	var person entity.Person
	err := json.Unmarshal(requestBody, &person)
	if err != nil {
		log.Fatalln(err)
	}
	db := Connector.Create(person)
	if db.Error != nil {
		log.Fatalln("-------------------> error while create = ", db.Error)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(person)
	if err != nil {
		log.Fatalln(err)
	}
}

func getPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person entity.Person
	Connector.First(&person, key)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(person)
	if err != nil {
		log.Fatalln(err)
	}
}

func updatePersonByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := io.ReadAll(r.Body)
	var person entity.Person
	err := json.Unmarshal(requestBody, &person)
	if err != nil {
		log.Fatalln(err)
	}
	Connector.Save(&person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(person)
	if err != nil {
		log.Fatalln(err)
	}
}

func deletPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person entity.Person
	id, _ := strconv.ParseInt(key, 10, 64)
	Connector.Where("id = ?", id).Delete(&person)
	w.WriteHeader(http.StatusNoContent)
}
