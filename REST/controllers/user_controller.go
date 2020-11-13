package controllers

import (
	"strconv"

	"github.com/arijitnayak92/taskAfford/REST/services"
	"github.com/arijitnayak92/taskAfford/REST/utils"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		apiError := &utils.APIError{
			Message:    "User Id shoudl be a number !",
			StatusCode: 400,
		}
		c.JSON(apiError.StatusCode, apiError)
		// jsonValue, _ := json.Marshal(apiError)
		// res.WriteHeader(apiError.StatusCode)
		// res.Write(jsonValue)
		return
	}
	user, apiError := services.UserService.GetUser(userId)
	if apiError != nil {
		// jsonValue, _ := json.Marshal(apiError)
		// res.WriteHeader(apiError.StatusCode)
		// res.Write(jsonValue)
		c.JSON(apiError.StatusCode, apiError)
		return
	}
	// jsonValue, _ := json.Marshal(user)
	c.JSON(200, user)
	//res.Write(200,user)
}
