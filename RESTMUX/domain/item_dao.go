package domain

import (
	"context"
	"fmt"
	"strconv"

	"github.com/arijitnayak92/taskAfford/RESTMUX/cache"
	"github.com/arijitnayak92/taskAfford/RESTMUX/utils"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	ItemDomain itemInterface
)

var itemCollection = db().Database("goAPI").Collection("items") // get collection "users" from db() which returns *mongo.Client

type itemInterface interface {
	AddItem(newItem *Item) (*Item, *utils.APIError)
	GetOne(itemID int64) (*Item, *utils.APIError)
	GetAll() ([]*Item, *utils.APIError)
	UpdateItem(itemID int64, newItem *Item) (*Item, *utils.APIError)
	DeleteItem(itemID int64) (*Item, *utils.APIError)
	Fibo(n int) (int, *utils.APIError)
}

func init() {
	ItemDomain = &itemStruct{}
	cache.InitializeRedis()
	cache.SetValue("0", "0")
	cache.SetValue("1", "1")
	cache.SetValue("2", "1")
}

type itemStruct struct {
	products []*Item
}

var lastNum = 3

func (c *itemStruct) Fibo(n int) (int, *utils.APIError) {
	if n >= 0 {
		if value, ok := cache.GetValue(fmt.Sprint(n)); ok {
			sValue, err := strconv.Atoi(value)
			if err != nil {
				return -1, &utils.APIError{
					Message:    "Operation Failed !",
					StatusCode: 422,
				}
			}
			return sValue, nil
		}
		for i := lastNum; i <= n; i++ {
			num1, err1 := cache.GetValue(fmt.Sprint(i - 1))
			num2, err2 := cache.GetValue(fmt.Sprint(i - 2))
			if !err1 || !err2 {
				return -1, &utils.APIError{
					Message:    "Operation Failed !",
					StatusCode: 422,
				}
			}
			num1Int, okD := strconv.Atoi(num1)
			num2Int, okP := strconv.Atoi(num2)
			if okD != nil || okP != nil {
				return -1, &utils.APIError{
					Message:    "Operation Failed !",
					StatusCode: 422,
				}
			}
			recent := num1Int + num2Int

			cache.SetValue(fmt.Sprint(i), fmt.Sprint(recent))
		}
		lastNum = n
		sendValue, _ := cache.GetValue(fmt.Sprint(n))
		sValue, _ := strconv.Atoi(sendValue)
		return sValue, nil
	}
	return -1, &utils.APIError{
		Message:    "Product Id Should be unique !",
		StatusCode: 406,
	}

}

func (c *itemStruct) AddItem(newItem *Item) (*Item, *utils.APIError) {
	found, _ := ItemDomain.GetOne(newItem.Id)

	if (found) != nil {
		return nil, &utils.APIError{
			Message:    "Product Id Should be unique !",
			StatusCode: 406,
		}
	}
	c.products = append(c.products, newItem)
	_, err := itemCollection.InsertOne(context.TODO(), newItem)
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Something went wrong !",
			StatusCode: 406,
		}
	}
	return &Item{}, nil
}

func (c *itemStruct) GetOne(itemID int64) (*Item, *utils.APIError) {
	var item *Item
	if err := itemCollection.FindOne(context.TODO(), bson.M{"id": itemID}).Decode(&item); err != nil {
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
	allData, err := itemCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Product Not Found !",
			StatusCode: 404,
		}
	}
	if errs := allData.All(context.TODO(), &items); errs != nil {
		return nil, &utils.APIError{
			Message:    "Product Not Found !",
			StatusCode: 404,
		}
	}
	c.products = items
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

	_, err := itemCollection.UpdateOne(
		context.TODO(),
		bson.M{"id": itemID},
		bson.D{
			{"$set", bson.M{"name": newItem.Name, "price": newItem.Price, "quantity": newItem.Quantity}},
		},
	)
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

	_, err := itemCollection.DeleteOne(context.TODO(), bson.M{"id": itemID})
	if err != nil {
		return nil, &utils.APIError{
			Message:    "Something went wrong !",
			StatusCode: 400,
		}
	}

	return &Item{}, nil
}
