package test

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"testing"
)

var ctx = context.Background()

func TestRedisClient(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})

	err := rdb.Set(ctx, "test1", "aaa", 0).Err()
	if err != nil {
		t.Fatal(err)
	}

	val, err := rdb.Get(ctx, "test1").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("test1:", val)
}
