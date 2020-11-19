package auth

import (
	"fmt"
	"net/http"

	"github.com/arijitnayak92/taskAfford/RESTMUXJWT/domain"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		err := domain.UserMethods.TokenValid(c.Request)
		fmt.Println("in auth")
		fmt.Println(err)
		if err != nil {
			fmt.Println("Got error")
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}
		accessDetails, _ := domain.UserMethods.ExtractTokenMetadata(c.Request)
		c.Set("accessDetails", accessDetails)
		c.Next()
	})
}
