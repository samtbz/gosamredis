package gosamredis

import (
	"context"

	"github.com/go-redis/redis"
)

type Samredis struct {
	rdb *redis.Client
	ctx context.Context
}

func Connect(Addr string) Samredis {
	var rdbStruct Samredis
	rdbStruct.rdb = redis.NewClient(&redis.Options{
		Addr:     Addr, //  "localhost:6379",
		Password: "",   // no password set
		DB:       0,    // use default DB
	})

	rdbStruct.ctx = context.Background()

	return rdbStruct

}

func (pointerToSamredis *Samredis) Set(key string, value string) {
	err := (*pointerToSamredis).rdb.Set((*pointerToSamredis).ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func (pointerToSamredis *Samredis) Get(key string) string {
	val, err := (*pointerToSamredis).rdb.Get((*pointerToSamredis).ctx, key).Result()
	if err != nil {
		panic(err)
	}

	return val
}
