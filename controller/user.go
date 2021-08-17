package controller

import (
	"fmt"
	"net/http"
	"portfolio/model"
	"portfolio/service"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	u := service.User{}

	users, err := u.GetUsers()
	if err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{})
	}

	c.JSONP(http.StatusOK, gin.H{
		"users": users,
	})
}

func CreateUser(c *gin.Context) {
	u := service.User{}

	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	username := user.Username
	password := fmt.Sprintf("%s", user.Password)

	if err := u.CreateUser(username, password); err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{"message": "internal server error"})
	} else {
		c.JSONP(http.StatusCreated, gin.H{
			"username": username,
		})
	}
}

func Login(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	username := user.Username
	password := string(user.Password)

	token, err := service.Login(username, password)
	if err != nil {
		c.JSONP(http.StatusInternalServerError, gin.H{"message": err})
	} else {
		c.JSONP(http.StatusCreated, gin.H{"token": token, "username": username})
	}
}

func Logout(c *gin.Context) {}
