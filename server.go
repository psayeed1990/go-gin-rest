package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//create GET, POST, PUT, DELETE REST API with gin
func main() {

	//check if GIN_MODE is set to 'release'
	//if it is, then set gin mode to release
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)

	}else{
		//check if .env file exists
		if _, err := os.Stat(".env"); err == nil {
			//load environment variables file with error checking
			if errs := godotenv.Load(); errs != nil {
				panic(errs)
			}
		
		}
	}


	

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	ConnectPostgres()

	// Create a group of routes
	v1 := router.Group("/v1")
	{
		v1.GET("/users", getAllUsers ) //get all users
		v1.GET("/users/:id", getUser) //get user by name
		v1.POST("/users", createUser) //create user
		v1.PUT("/users/:id", updateUser) //update user
		v1.DELETE("/users/:id", deleteUser)	//delete user
	} 

	// PORT environment variable was defined.
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}
	router.Run(":"+PORT)
}