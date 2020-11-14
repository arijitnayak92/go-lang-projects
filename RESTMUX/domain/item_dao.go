package domain

import (
	"context"
	"fmt"

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
}

func init() {
	ItemDomain = &itemStruct{}
	cache.InitializeRedis()
}

type itemStruct struct {
	products []*Item
}

var lastNum int = 3

func (c *itemStruct) Fibo(n int) (int, *utils.APIError) {
	if n >= 0 {
		if value, ok := cache.GetValue("FiboN", n); ok {
			return value, nil
		} else {
			for i := lastNum; i <= n; i++ {
				redis := new(FiboStruct)
				recent = cache.GetValue("FiboN", i-1) + cache.GetValue("FiboN", i-2)
				redis.ForID = string(i)
				redis.value = string(recent)
				cache.SetValue("FiboN", redis)
			}
			lastNum = n
			return localCache[n], nil
		}
	} else {
		return -1, &utils.APIError{
			Message:    "Product Id Should be unique !",
			StatusCode: 406,
		}
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
