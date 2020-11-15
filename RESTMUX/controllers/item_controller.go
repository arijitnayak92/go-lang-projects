package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/arijitnayak92/taskAfford/RESTMUX/domain"
	"github.com/arijitnayak92/taskAfford/RESTMUX/services"
	"github.com/arijitnayak92/taskAfford/RESTMUX/utils"
	"github.com/gorilla/mux"
)

func GetNthFibonacii(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id := params["number"]
	intID, err := strconv.Atoi(id)
	if err != nil {
		apiError := &utils.APIError{
			Message:    "User Id shoudl be a number !",
			StatusCode: 400,
		}
		jsonValue, _ := json.Marshal(apiError)
		res.WriteHeader(apiError.StatusCode)
		res.Write(jsonValue)
		return
	}
	value, errors := services.ItemServicePublic.Fibonacii(intID)
	if errors != nil {
		jsonValue, _ := json.Marshal(errors)
		res.WriteHeader(errors.StatusCode)
		res.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(value)
	res.Write(jsonValue)
}

func GetOneProduct(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id := params["item_id"]
	itemID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		apiError := &utils.APIError{
			Message:    "User Id shoudl be a number !",
			StatusCode: 400,
		}
		jsonValue, _ := json.Marshal(apiError)
		res.WriteHeader(apiError.StatusCode)
		res.Write(jsonValue)
		return
	}
	item, errors := services.ItemServicePublic.GetOneItem(itemID)
	if errors != nil {
		jsonValue, _ := json.Marshal(errors)
		res.WriteHeader(errors.StatusCode)
		res.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(item)
	res.Write(jsonValue)
}

func GetAllProduct(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	items, errors := services.ItemServicePublic.GetAllItem()
	if errors != nil {
		jsonValue, _ := json.Marshal(errors)
		res.WriteHeader(errors.StatusCode)
		res.Write(jsonValue)
		return
	}
	jsonValue, _ := json.Marshal(items)
	res.Write(jsonValue)
}

func DeleteOneProduct(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id := params["item_id"]
	itemID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		apiError := &utils.APIError{
			Message:    "User Id shoudl be a number !",
			StatusCode: 400,
		}
		jsonValue, _ := json.Marshal(apiError)
		res.WriteHeader(apiError.StatusCode)
		res.Write(jsonValue)
		return
	}

	_, errors := services.ItemServicePublic.DeleteItem(itemID)
	if errors != nil {
		jsonValue, _ := json.Marshal(errors)
		res.WriteHeader(errors.StatusCode)
		res.Write(jsonValue)
		return
	}
	res.Write([]byte("Item Deleted !"))
}

func UpdateOneProduct(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	id := params["item_id"]
	itemID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		apiError := &utils.APIError{
			Message:    "User Id shoudl be a number !",
			StatusCode: 400,
		}
		jsonValue, _ := json.Marshal(apiError)
		res.WriteHeader(apiError.StatusCode)
		res.Write(jsonValue)
		return
	}

	newItem := new(domain.Item)
	errNew := json.NewDecoder(req.Body).Decode(&newItem)
	if errNew != nil {
		apiError := &utils.APIError{
			Message:    "Insufficient Data !",
			StatusCode: 406,
		}
		jsonValue, _ := json.Marshal(apiError)
		res.WriteHeader(apiError.StatusCode)
		res.Write(jsonValue)
		return
	}

	_, errors := services.ItemServicePublic.UpdateItem(itemID, newItem)
	if errors != nil {
		jsonValue, _ := json.Marshal(errors)
		res.WriteHeader(errors.StatusCode)
		res.Write(jsonValue)
		return
	}
	res.Write([]byte("Item Updated  !"))
}

func AddProduct(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	newItem := new(domain.Item)
	err := json.NewDecoder(req.Body).Decode(&newItem)
	if err != nil {
		apiError := &utils.APIError{
			Message:    "Insufficient Data !",
			StatusCode: 406,
		}
		jsonValue, _ := json.Marshal(apiError)
		res.WriteHeader(apiError.StatusCode)
		res.Write(jsonValue)
	}
	_, errors := services.ItemServicePublic.AddItem(newItem)
	if errors != nil {
		jsonValue, _ := json.Marshal(errors)
		res.WriteHeader(errors.StatusCode)
		res.Write(jsonValue)
		return
	}
	res.Write([]byte("Successfully added !"))
}
