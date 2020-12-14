package auth

import (
	"fmt"
	"net/http"

	"github.com/arijitnayak92/taskAfford/Fruit/utils"
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		u := utils.NewUtil()
		err := u.TokenValid(c.Request)
		fmt.Println("in auth")
		fmt.Println(err)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		accessDetails, _ := u.ExtractTokenMetadata(c.Request)
		c.Set("accessDetails", accessDetails)
		c.Next()
	})
}
