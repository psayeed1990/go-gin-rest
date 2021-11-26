package main

//import postgres gorm
// Language: go
// Path: db.go
import (
	"os"

	"github.com/jinzhu/gorm"

	// _ "github.com/jinzhu/gorm/dialects/sqlite"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

// func ConnectDB() {
// 	database, err := gorm.Open("sqlite3", "test.db")
  
// 	if err != nil {
// 	  panic("Failed to connect to database!")
// 	}
  
// 	database.AutoMigrate(&User{})
  
// 	DB = database
// }


func ConnectPostgres() {
	//get DB url from environment
	var DB_NAME = os.Getenv("POSTGRES_DB")
	var HOST = os.Getenv("POSTGRES_HOST")
	var DB_USER = os.Getenv("POSTGRES_USER")
	var DB_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	
	dsn := "host=" + HOST + " user=" + DB_USER + " dbname=" + DB_NAME + " password=" + DB_PASSWORD + " sslmode=disable"

	//connect postgres
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect database")
	}
	
	db.AutoMigrate(&User{})
	DB = db
	//defer db.Close()
}