package main

import (
	godotenvangel "DT/Env"
	models "DT/Models"
	"DT/Router"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenvangel.Load(); err != nil {
		log.Print("No .env file found")
	}
	os.Setenv("PWD", "abc#$")
}

func main() {
	models.ConnectDataBase() // new
	Router.Router()
}
