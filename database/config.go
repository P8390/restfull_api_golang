package database

import (
	"fmt"
	"log"
)

//Config to maintain DB configuration properties
type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

// GetConnectionString ...
var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.User, config.Password, config.ServerName, config.DB)
	log.Println("connectionString is - " + connectionString)
	return connectionString
}
