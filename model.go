package main

type User struct {
	ID uint `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Email string  `json:"email"`
	Password string `json:"password"`
}

type UserInput struct {
	Name string `json:"name" binding:"required"`
	Email string  `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateInput struct{
	Name string `json:"name"`
	Email string  `json:"email"`
	Password string `json:"password"`
}