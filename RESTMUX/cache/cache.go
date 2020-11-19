package cache

import (
	"encoding/json"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func InitializeRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:       "localhost:6379",
		PoolSize:   100,
		MaxRetries: 2,
		Password:   "",
		DB:         0,
	})

	ping, err := redisClient.Ping().Result()
	if err == nil && len(ping) > 0 {
		println("Connected to Redis")
	} else {
		println("Redis Connection Failed")
	}
}

func GetValue(key string) (string, bool) {
	var deserializedValue interface{}
	serializedValue, err := redisClient.Get(key).Result()
	if err != nil {
		return "", false
	}
	json.Unmarshal([]byte(serializedValue), &deserializedValue)
	return deserializedValue.(string), true
}

func SetValue(key string, value interface{}) (bool, error) {
	serializedValue, _ := json.Marshal(value)
	err := redisClient.Set(key, string(serializedValue), 0).Err()
	return true, err
}

func DelKey(key string) error {
	return redisClient.Del(key).Err()
}
