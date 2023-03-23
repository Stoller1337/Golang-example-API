package internal

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		config.User, config.Password, config.ServerName, config.DB)
	return connectionString
}

var Connector *gorm.DB

// Connect creates MySQL connection
func Connect(connectionString string) error {
	_, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	} else {
		log.Println("Connection to db was successful")
	}
	return nil
}
