package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// NewRedisClient is placeholder
//
//	func NewRedisClient() *redis.Client {
//		db, _ := strconv.Atoi("0")
//		return redis.NewClient(&redis.Options{
//			Addr:     fmt.Sprintf("%s:%s", "178.128.125.218", "16399"),
//			Password: "a2Ozhm243oj4yl",
//			DB:       db,
//		})
//	}
func NewRedisClient() *redis.Client {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})
}

// Get is placeholder
func Get(rdb *redis.Client, key string) string {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println(key + " does not exist")
	} else if err != nil {
		fmt.Println(err)
	}
	return string(val)
}

// Delete is placeholder
func Delete(rdb *redis.Client, key string) error {
	return rdb.Del(ctx, key).Err()
}

// SetString is placeholder
func SetString(rdb *redis.Client, key string, value string, expiration time.Duration) error {
	return rdb.Set(ctx, key, value, expiration).Err()
}

// SetJSON is placeholder
func SetJSON(rdb *redis.Client, key string, value interface{}, expiration time.Duration) error {
	jsonStr, _ := json.Marshal(value)
	return rdb.Set(ctx, key, jsonStr, expiration).Err()
}
