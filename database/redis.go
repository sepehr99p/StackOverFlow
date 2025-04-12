package database

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0, // use default DB
	})

	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}
}

func CacheQuestion(ctx context.Context, key string, value interface{}) error {
	return RedisClient.Set(ctx, key, value, 5*time.Minute).Err()
}

func GetCachedQuestion(ctx context.Context, key string) (string, error) {
	return RedisClient.Get(ctx, key).Result()
}

func CacheUserToken(ctx context.Context, key string, value interface{}) error {
	return RedisClient.Set(ctx, key, value, 24*time.Hour).Err()
}

func GetCachedUserToken(ctx context.Context, key string) (string, error) {
	return RedisClient.Get(ctx, key).Result()
}

func DeleteCachedToken(ctx context.Context, key string) error {
	return RedisClient.Del(ctx, key).Err()
}
