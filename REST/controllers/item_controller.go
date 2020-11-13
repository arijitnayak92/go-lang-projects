package controllers

import (
	"strconv"

	"github.com/arijitnayak92/taskAfford/REST/domain"
	"github.com/arijitnayak92/taskAfford/REST/services"
	"github.com/arijitnayak92/taskAfford/REST/utils"
	"github.com/gin-gonic/gin"
)

func GetOneProduct(c *gin.Context) {
	itemID, err := strconv.ParseInt(c.Param("item_id"), 10, 64)
	if err != nil {
		apiError := &utils.APIError{
			Message:    "User Id shoudl be a number !",
			StatusCode: 400,
		}
		c.JSON(apiError.StatusCode, apiError)
		return
	}
	item, errors := services.ItemServicePublic.GetOneItem(itemID)
	if errors != nil {
		c.JSON(errors.StatusCode, errors)
		return
	}
	c.JSON(200, item)
}

func GetAllProduct(c *gin.Context) {
	items, errors := services.ItemServicePublic.GetAllItem()
	if errors != nil {
		c.JSON(errors.StatusCode, errors)
		return
	}
	c.JSON(200, items)
}

func DeleteOneProduct(c *gin.Context) {
	itemID, err := strconv.ParseInt(c.Param("item_id"), 10, 64)
	if err != nil {
		apiError := &utils.APIError{
			Message:    "User Id shoudl be a number !",
			StatusCode: 400,
		}
		c.JSON(apiError.StatusCode, apiError)
		return
	}

	_, errors := services.ItemServicePublic.DeleteItem(itemID)
	if errors != nil {
		c.JSON(errors.StatusCode, errors)
		return
	}
	c.JSON(200, "Item Deleted Successfully !")
}

func UpdateOneProduct(c *gin.Context) {
	itemID, err := strconv.ParseInt(c.Param("item_id"), 10, 64)
	if err != nil {
		apiError := &utils.APIError{
			Message:    "User Id shoudl be a number !",
			StatusCode: 400,
		}
		c.JSON(apiError.StatusCode, apiError)
		return
	}

	newItem := new(domain.Item)
	err_new := c.BindJSON(&newItem)
	if err_new != nil {
		apiError := &utils.APIError{
			Message:    "Insufficient Data !",
			StatusCode: 406,
		}
		c.JSON(apiError.StatusCode, apiError)
		return
	}

	_, errors := services.ItemServicePublic.UpdateItem(itemID, newItem)
	if errors != nil {
		c.JSON(errors.StatusCode, errors)
		return
	}
	c.JSON(200, "Item Updated Successfully !")
}

func AddProduct(c *gin.Context) {
	newItem := new(domain.Item)
	err := c.BindJSON(&newItem)
	if err != nil {
		apiError := &utils.APIError{
			Message:    "Insufficient Data !",
			StatusCode: 406,
		}
		c.JSON(apiError.StatusCode, apiError)
		return
	}
	_, errors := services.ItemServicePublic.AddItem(newItem)
	if errors != nil {
		c.JSON(errors.StatusCode, errors)
		return
	}
	c.JSON(200, "Successfully Added !")
}
