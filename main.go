package main

import (
	"awesomeProject2/internal"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main() {
	config :=
		internal.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "4512",
			DB:         "learning",
		}

	connectionString := internal.GetConnectionString(config)
	if err := internal.Connect(connectionString); err != nil {
		log.Fatal(err)
	}
	internal.RunServer()
}
