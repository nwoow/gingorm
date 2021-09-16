package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

//DB ...
var DB *gorm.DB

//ConnectDataBase ...
func ConnectDataBase() {

	// github.com/mattn/go-sqlite3
	// db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	DBHOST := os.Getenv("DB_HOST")
	DBUSERNAME := os.Getenv("DB_USERNAME")
	DBPASSWORD := os.Getenv("DB_PASSWORD")
	// var databasePass string
	// databasePass = os.Getenv("DB_PASSWORD")
	// PASSWORD := "ang31#$"
	DBNAME := os.Getenv("DB_NAME")
	sqlURL := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", DBHOST, DBUSERNAME, DBPASSWORD, DBNAME, "5432", "disable")
	// // fmt.Println("Db password:", databasePass)
	// // fmt.Println("Db password1:", os.Getenv("PWD"))
	// // os.Exit(1)
	database, err := gorm.Open(postgres.Open(sqlURL), &gorm.Config{})

	// // database, err := sql.Open("mssql", sqlURL)
	fmt.Println("database", sqlURL)

	if err != nil {
		panic("Failed to connect to database!")
	}
	// database.AutoMigrate(&Logger{}, &Logconfigure{},&Book{})
	DB = database
}
