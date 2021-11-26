package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//create GET, POST, PUT, DELETE REST API with gin
func main() {
	//load dotenv
	err := godotenv.Load()

	//check error to load environment variable
	if err != nil {
		panic(err)
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