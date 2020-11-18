package controllers

import (
	"fmt"
	"net/http"

	"github.com/arijitnayak92/taskAfford/REST/domain"
	"github.com/arijitnayak92/taskAfford/REST/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	fmt.Println(c.Request.Header.Get("Authorization"))
	var u *domain.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	fmt.Println(u)
	if u.Username == "" || u.Password == "" {
		c.JSON(406, "Enter all the details !")
	}

	user, apiError := services.UserService.CreateUser(u)
	if apiError != nil {
		c.JSON(apiError.StatusCode, apiError)
		return
	}
	// jsonValue, _ := json.Marshal(user)
	c.JSON(200, user)
	//res.Write(200,user)
}

func Login(c *gin.Context) {
	var u *domain.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	fmt.Println(u)
	if u.Username == "" || u.Password == "" {
		c.JSON(406, "Enter all the details !")
	}

	user, apiError := services.UserService.Login(u)
	if apiError != nil {
		c.JSON(apiError.StatusCode, apiError)
		return
	}
	// jsonValue, _ := json.Marshal(user)
	c.JSON(200, user)
	//res.Write(200,user)
}

func RefreshToken(c *gin.Context) {
	user, apiError := services.UserService.RefreshToken(c)
	if apiError != nil {
		c.JSON(apiError.StatusCode, apiError)
		return
	}
	// jsonValue, _ := json.Marshal(user)
	c.JSON(200, user)
	//res.Write(200,user)
}
