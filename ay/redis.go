package ay

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var (
	Redis    *redis.Client
	RedisNil = redis.Nil
)

func InitializeRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     Yaml.GetString("redis.addr"),
		Password: Yaml.GetString("redis.password"),
		DB:       Yaml.GetInt("redis.db"),
		PoolSize: Yaml.GetInt("redis.pool_size"),
	})
	err := Redis.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
}
