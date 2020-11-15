package services

import (
	"github.com/arijitnayak92/taskAfford/RESTMUX/domain"
	"github.com/arijitnayak92/taskAfford/RESTMUX/utils"
)

var (
	ItemServicePublic itemServicesInterface
)

type itemServicesInterface interface {
	AddItem(newItem *domain.Item) (*domain.Item, *utils.APIError)
	GetOneItem(itemID int64) (*domain.Item, *utils.APIError)
	GetAllItem() ([]*domain.Item, *utils.APIError)
	UpdateItem(itemID int64, newItem *domain.Item) (*domain.Item, *utils.APIError)
	DeleteItem(itemID int64) (*domain.Item, *utils.APIError)
	Fibonacii(number int) (int, *utils.APIError)
}

func init() {
	ItemServicePublic = &itemServices{}
}

type itemServices struct{}

func (c *itemServices) Fibonacii(number int) (int, *utils.APIError) {
	return domain.ItemDomain.Fibo(number)
}

func (c *itemServices) AddItem(newItem *domain.Item) (*domain.Item, *utils.APIError) {
	return domain.ItemDomain.AddItem(newItem)
}

func (c *itemServices) GetOneItem(itemID int64) (*domain.Item, *utils.APIError) {
	return domain.ItemDomain.GetOne(itemID)
}

func (c *itemServices) GetAllItem() ([]*domain.Item, *utils.APIError) {
	return domain.ItemDomain.GetAll()
}

func (c *itemServices) UpdateItem(itemID int64, newItem *domain.Item) (*domain.Item, *utils.APIError) {
	return domain.ItemDomain.UpdateItem(itemID, newItem)
}

func (c *itemServices) DeleteItem(itemID int64) (*domain.Item, *utils.APIError) {
	return domain.ItemDomain.DeleteItem(itemID)
}
