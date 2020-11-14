package main

import "github.com/arijitnayak92/taskAfford/REST-FIBO-REDIS/cache"

func main() {

	cache.InitializeRedis()

	println("Setting Testkey -> TestValue")
	cache.SetValue("TestKey", "TestValue")

	println("Getting TestKey")
	value, err := cache.GetValue("TestKey")

	if err == nil {
		println("Value Returned : " + value.(string))
	} else {
		println("Getting Value Failed with error : " + err.Error())
	}

}

/*

package cache

import (
    "context"
		"fmt"
		"encoding/json"
    "github.com/go-redis/redis/v8"
)


type Author struct {
	for string `json:"name"`
	value  int    `json:"age"`
}

func Redis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(Author{for: "Arijit", value: 22})
	if err != nil {
		fmt.Println(err)
	}

	err = client.Set(context.Background(), "id1234", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get(context.Background(), "id1234").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}

*/
