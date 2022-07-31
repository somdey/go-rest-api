package controllers

import (
	"fmt"
	"go_web_app/users"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Summary CreateUser
// @Description create a new user in db
// @Accept  json
// @Produce  json
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} object
// @Failure 404 {object} object
// @Failure 500 {object} object
// @Param body body model.User true "Create user"
// @Router /users [post]
func CreateUser(c *gin.Context) {
	body := c.Request.Body
	x, _ := ioutil.ReadAll(body)

	fmt.Printf("%s input\n", string(x))
	res := users.CreateUser(x)
	fmt.Printf("%s mongo\n", string(res))
	c.JSON(200, gin.H{
		"message": string(x),
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "user has been updated",
	})
}

func GetUser(c *gin.Context) {
	fmt.Printf("Get user method called.")
	c.JSON(200, gin.H{
		"message": "Get a single user",
	})
}

func GetUsers(c *gin.Context) {
	d := users.GetUsers()
	fmt.Print(string(d))
	c.JSON(200, gin.H{"message": d})
}

func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete a user",
	})
}
