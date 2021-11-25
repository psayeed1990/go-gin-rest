package main

//import gin
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//get all users
func getAllUsers(c *gin.Context) {

	var users []User
	DB.Find(&users)

	c.JSON(200, gin.H{
		"Users": users,
	})
}

//getUser
func getUser(c *gin.Context) {
	var user User

	//find user by id
	if err := DB.First(&user, c.Param("id")).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		
		return
	}

	c.JSON(200, gin.H{
		"User": user,
	})
}

//createUser
func createUser(c *gin.Context) {
	//create user input
	var input UserInput
	//check error of the input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	//create user
	user := User{
		Name: input.Name,
		Email: input.Email,
		Password: input.Password,
	}
	
	//save user
	DB.Create(&user)

	c.JSON(200, gin.H{
		"New User": user,
	})
}

//updateUser
func updateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "update user",
	})
}

//deleteUser
func deleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete user",
	})
}
