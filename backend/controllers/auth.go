package controllers

import (
	"fmt"
	"github.com/Rageon787/file-hosting-service/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {

	var user models.User

	// Copy the JSON payload into the newUser?
	if err := c.BindJSON(&user); err != nil {
		return
	}

	_ = fmt.Sprintf("Email: %s Name: %s Password: %s", user.Email, user.Name, user.Password)

	// search if the user exists in the DB
	// compare provided password with password in DB?

	c.IndentedJSON(http.StatusCreated, user)
}

func Signup(c *gin.Context) {
	var newUser models.User

	// Copy the JSON payload into the newUser?
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	_ = fmt.Sprintf("Email: %s Name: %s Password: %s", newUser.Email, newUser.Name, newUser.Password)

	// TODO: Insert the new user into the database

	c.IndentedJSON(http.StatusCreated, newUser)

}

func Logout(c *gin.Context) {

}
