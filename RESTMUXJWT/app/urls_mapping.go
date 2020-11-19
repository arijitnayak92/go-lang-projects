package app

import (
	"github.com/arijitnayak92/taskAfford/RESTMUXJWT/auth"
	"github.com/arijitnayak92/taskAfford/RESTMUXJWT/controllers"
	"github.com/gin-gonic/gin"
)

func response(c *gin.Context) {
	c.JSON(200, "Backend Connected !")
}

func Routes() {
	router.GET("/", response)
	router.POST("/login", controllers.Login)
	router.POST("/logout", controllers.Logout)
	router.POST("/create", controllers.CreateUser)
	router.POST("/refreshToken", controllers.RefreshToken)
	router.GET("/getOneItem/:item_id", auth.TokenAuthMiddleware(), controllers.GetOneProduct)
	router.GET("/getAllItem", auth.TokenAuthMiddleware(), controllers.GetAllProduct)
	router.POST("/addItem", auth.TokenAuthMiddleware(), controllers.AddProduct)
	router.PUT("/updateItem/:item_id", auth.TokenAuthMiddleware(), controllers.UpdateOneProduct)
	router.DELETE("/deleteItem/:item_id", auth.TokenAuthMiddleware(), controllers.DeleteOneProduct)
}
