package database

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func RedisClientSet(x, y int, value []int) {
	valueRadis := fmt.Sprintf("%v", value)
	res := x + y
	data := []byte(fmt.Sprintf("%v", res))
	hash := sha256.Sum256(data)
	keyHash := fmt.Sprintf("%x", hash[:])

	Rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := Rdb.Set(ctx, keyHash, valueRadis, 0).Err()
	if err != nil {
		panic(err)
	}

}

func RedisClientGet(x, y int) (string, error) {

	res := x + y
	data := []byte(fmt.Sprintf("%v", res))
	hash := sha256.Sum256(data)
	keyHash := fmt.Sprintf("%x", hash[:])

	Rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	redisValue, err := Rdb.Get(ctx, keyHash).Result()

	if err != nil {
		return redisValue, err
	}
	return redisValue, nil

}
