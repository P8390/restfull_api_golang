package main

import (
	"log"
	"net/http"
	"rest-go-demo/database"

	"github.com/gorilla/mux"

	"rest-go-demo/controllers"

	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect

	"rest-go-demo/entity"
)

func main() {
	initDB()
	log.Println("Starting the HTTP Server on Port 8090")

	router := mux.NewRouter()
	initialiseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initialiseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/get", controllers.GetAllPerson).Methods("GET")
	router.HandleFunc("/get/{id}", controllers.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletePersonByID).Methods("DELETE")
}

func initDB() {
	config := database.Config{
		ServerName: "localhost:3306",
		User:       "root",
		Password:   "Winter01!",
		DB:         "smsmagic",
	}
	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Person{})

}
