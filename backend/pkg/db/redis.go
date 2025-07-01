package db

import (
    "github.com/go-redis/redis/v8"
    "wowtoken-api/internal/config"
)

func GetRedisClient() *redis.Client {
    cfg, _ := config.LoadConfig()
    return redis.NewClient(&redis.Options{
        Addr: cfg.RedisAddr,
    })
}