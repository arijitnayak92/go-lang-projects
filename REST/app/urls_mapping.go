package app

import (
	"github.com/arijitnayak92/taskAfford/REST/controllers"
	"github.com/gin-gonic/gin"
)

func response(c *gin.Context) {
	c.JSON(200, "Backend Connected !")
}

func Routes() {
	router.GET("/", response)
	router.GET("/users", controllers.GetUser)
	router.GET("/getOneItem/:item_id", controllers.GetOneProduct)
	router.GET("/getAllItem", controllers.GetAllProduct)
	router.POST("/addItem", controllers.AddProduct)
	router.PUT("/updateItem/:item_id", controllers.UpdateOneProduct)
	router.DELETE("/deleteItem/:item_id", controllers.DeleteOneProduct)
}
