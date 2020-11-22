package domain

import (
	"context"
	"fmt"

	"github.com/arijitnayak92/taskAfford/RESTTODO/utils"
)

var (
	ItemDomain itemInterface
	pg *pg.DB
)

type itemInterface interface {
	AddItem(newItem *Item) (uint64, *utils.APIError)
	GetOne(itemID int64) (*Item, *utils.APIError)
	GetAll() ([]*Item, *utils.APIError)
	UpdateItem(,itemID int64, newItem *Item) (*Item, *utils.APIError)
	DeleteItem(itemID int64) (*Item, *utils.APIError)
}

func init() {
	ItemDomain = &itemStruct{}
}

type itemStruct struct {
	products []*Item
}

func (c *itemStruct) AddItem(newItem *Item) (uint64, *utils.APIError) {

	insertedID, err := pg.Model(newItem).Returning("id").Insert()
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Error Occoured while inserting the data !",
			StatusCode: 422,
		}
	}
	return insertedID, nil
}

func (c *itemStruct) GetOne(itemID int64) (*Item, *utils.APIError) {
	var item *Item
	if err := pg.Model(item).Where("item.id=?", item.ID).Select(); err != nil {
		fmt.Println(err)
		return nil, &utils.APIError{
			Message:    "Product Not Found !",
			StatusCode: 404,
		}
	}
	return item, nil
}

func (c *itemStruct) GetAll() ([]*Item, *utils.APIError) {
	var items []*Item
	err := pg.Model(items).Select()
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Product Not Found !",
			StatusCode: 404,
		}
	}
	return items, nil
}

func (c *itemStruct) UpdateItem(itemID int64, newItem *Item) (*Item, *utils.APIError) {
	_, errors := ItemDomain.GetOne(itemID)
	if errors != nil {
		return nil, &utils.APIError{
			Message:    errors.Message,
			StatusCode: errors.StatusCode,
		}
	}

	_, err := pg.Model(newItem).Column("title", "description").Where("id = ?0", itemID).Update()

	if err != nil {
		return nil, &utils.APIError{
			Message:    "Error in processing data !",
			StatusCode: 422,
		}
	}

	return &Item{}, nil
}

func (c *itemStruct) DeleteItem(itemID int64) (*Item, *utils.APIError) {
	_, errors := ItemDomain.GetOne(itemID)
	if errors != nil {
		return nil, &utils.APIError{
			Message:    errors.Message,
			StatusCode: errors.StatusCode,
		}
	}
	var item *Item
	_, err := pg.Model(item).Where("id = ?0", itemID).Delete()
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Error in processing data !",
			StatusCode: 422,
		}
	}

	return &Item{}, nil
}
