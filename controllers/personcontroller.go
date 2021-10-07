package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"rest-go-demo/database"
	"rest-go-demo/entity"

	"github.com/gorilla/mux"
)

//GetAllPerson get all person data
func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside GetAllPerson")
	var persons []entity.Person
	database.Connector.Find(&persons)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(persons)
}

//GetPersonByID returns person with specific ID
func GetPersonByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside GetPersonByIDgo")
	vars := mux.Vars(r)
	key := vars["id"]
	var person entity.Person
	database.Connector.First(&person, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

//CreatePerson creates person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside CreatePerson")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Person

	json.Unmarshal(reqBody, &person)
	database.Connector.Create(person)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

//UpdatePersonByID updates person with respective ID
func UpdatePersonByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside UpdatePersonByID")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Person
	json.Unmarshal(reqBody, &person)
	database.Connector.Save(&person)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

//DeletePersonByID delete's person with specific ID
func DeletePersonByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Inside DeletePersonByID")
	vars := mux.Vars(r)
	key := vars["id"]

	var person entity.Person
	database.Connector.Where("id = ?", key).Delete(&person)
	w.WriteHeader(http.StatusNoContent)
}
