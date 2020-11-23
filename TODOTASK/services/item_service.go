package services

import (
	"github.com/arijitnayak92/taskAfford/RESTMUXJWT/domain"
	"github.com/arijitnayak92/taskAfford/RESTMUXJWT/utils"
)

var (
	ItemServicePublic itemServicesInterface
)

type itemServicesInterface interface {
	AddItem(newItem *domain.Item) (*domain.Item, *utils.APIError)
	GetOneItem(itemID uint64) (*domain.Item, *utils.APIError)
	GetAllItem() ([]*domain.Item, *utils.APIError)
	UpdateItem(itemID uint64, newItem *domain.Item) (bool, *utils.APIError)
	DeleteItem(itemID uint64) (bool, *utils.APIError)
}

func init() {
	ItemServicePublic = &itemServices{}
}

type itemServices struct{}

func (c *itemServices) AddItem(newItem *domain.Item) (*domain.Item, *utils.APIError) {
	return domain.ItemDomain.AddItem(newItem)
}

func (c *itemServices) GetOneItem(itemID int64) (*domain.Item, *utils.APIError) {
	return domain.ItemDomain.GetOne(itemID)
}

func (c *itemServices) GetAllItem() ([]*domain.Item, *utils.APIError) {
	return domain.ItemDomain.GetAll()
}

func (c *itemServices) UpdateItem(itemID uint64, newItem *domain.Item) (bool, *utils.APIError) {
	return domain.ItemDomain.UpdateItem(itemID, newItem)
}

func (c *itemServices) DeleteItem(itemID uint64) (bool, *utils.APIError) {
	return domain.ItemDomain.DeleteItem(itemID)
}
