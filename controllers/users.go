package controllers

import "github.com/gin-gonic/gin"

func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user has been created",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user has been updated",
	})
}

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get a single user",
	})
}

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get all users",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete a user",
	})
}
