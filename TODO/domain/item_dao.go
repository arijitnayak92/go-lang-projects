package domain

import (
	"database/sql"
	"fmt"

	"github.com/arijitnayak92/taskAfford/TODO/utils"

	_ "github.com/lib/pq"
)

var (
	ItemDomain itemInterface
	dbIns      *sql.DB
)

type itemInterface interface {
	AddItem(newItem *Item) (*Item, *utils.APIError)
	GetOne(itemID int64) (*Item, *utils.APIError)
	GetAll() ([]*Item, *utils.APIError)
	UpdateItem(itemID int64, newItem *Item) (*Item, *utils.APIError)
	DeleteItem(itemID int64) (*Item, *utils.APIError)
}

func init() {
	fmt.Println(dbIns)
	ItemDomain = &itemStruct{}
}

type itemStruct struct {
	products []*Item
}

func (c *itemStruct) AddItem(newItem *Item) (*Item, *utils.APIError) {
	found, _ := ItemDomain.GetOne(newItem.Id)

	if (found) != nil {
		return nil, &utils.APIError{
			Message:    "Product Id Should be unique !",
			StatusCode: 406,
		}
	}
	statement, err := dbIns.Prepare("INSERT INTO todo(title, description, status) VALUES($1, $2, $3);")
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Something went wrong !",
			StatusCode: 406,
		}
	}
	statement.QueryRow(newItem)
	return &Item{}, nil
}

func (c *itemStruct) GetOne(itemID int64) (*Item, *utils.APIError) {
	statement, err := dbIns.Prepare("SELECT * FROM todo WHERE id= $1;")
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Error in DBInstances !",
			StatusCode: 422,
		}
	}
	row := statement.QueryRow(statement, itemID)
	var taskID int64
	var title string
	var description string
	var status bool

	switch errs := row.Scan(&taskID, &title, &description, &status); errs {
	case sql.ErrNoRows:
		return nil, &utils.APIError{
			Message:    "Product Not Found !",
			StatusCode: 404,
		}
	case nil:
		return &Item{
			Id:          itemID,
			Title:       title,
			Description: description,
			Status:      status,
		}, nil
	default:
		return nil, &utils.APIError{
			Message:    "Product Not Found !",
			StatusCode: 404,
		}
	}
}

func (c *itemStruct) GetAll() ([]*Item, *utils.APIError) {
	var items []*Item

	statement, err := dbIns.Prepare("SELECT * FROM todo;")
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Error in DBInstances !",
			StatusCode: 422,
		}
	}
	rows, errs := statement.Query(statement)
	if errs != nil {
		return nil, &utils.APIError{
			Message:    "Products Not Found !",
			StatusCode: 404,
		}
	}
	var taskID int64
	var title string
	var description string
	var status bool
	i := 0
	for rows.Next() {
		i++
		rows.Scan(&taskID, &title, &description, &status)
		items = append(items, &Item{
			Id:          taskID,
			Title:       title,
			Description: description,
			Status:      status,
		})
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

	statement, err := dbIns.Prepare("UPDATE todo SET title = $1,description = $2 WHERE id = $3;")
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Error in processing data !",
			StatusCode: 422,
		}
	}
	statement.QueryRow(newItem, itemID)

	if err != nil {
		return nil, &utils.APIError{
			Message:    "Something went wrong !",
			StatusCode: 400,
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
	statement, err := dbIns.Prepare("DELETE FROM todo WHERE id = $1;")
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Error in processing data !",
			StatusCode: 422,
		}
	}
	statement.QueryRow(itemID)

	return &Item{}, nil
}
