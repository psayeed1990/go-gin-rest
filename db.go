package main

//import postgres gorm
// Language: go
// Path: db.go
import (
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
	// var DATABASE_URL = os.Getenv("DATABASE_URL")
	// log.Println("db"+DATABASE_URL)
	//connect postgres
	db, err := gorm.Open("postgres", "host=localhost user=sayeed password=1990 dbname=diesel_demo sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	
	db.AutoMigrate(&User{})
	DB = db
	//defer db.Close()
}