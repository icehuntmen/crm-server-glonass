package cache

import (
	"context"
	"crm-glonass/config"
	"crm-glonass/pkg/logging"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var logcod = logging.NewLogger(config.GetConfig())
var redisClient *redis.Client

func InitRedis(cfg *config.Config, ctx context.Context) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password:           cfg.Redis.Password,
		DB:                 cfg.Redis.Db,
		DialTimeout:        cfg.Redis.DialTimeout,
		ReadTimeout:        cfg.Redis.ReadTimeout,
		WriteTimeout:       cfg.Redis.WriteTimeout,
		PoolSize:           cfg.Redis.PoolSize,
		PoolTimeout:        cfg.Redis.PoolTimeout,
		IdleTimeout:        500 * time.Millisecond,
		IdleCheckFrequency: cfg.Redis.IdleCheckFrequency * time.Millisecond,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}
	logcod.Info(logging.Redis, logging.Connection, "Redis cache init", nil)
	return nil
}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}

func Set[T any](ctx context.Context, c *redis.Client, key string, value T, duration time.Duration) error {

	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.Set(ctx, key, v, duration).Err()
}

func Get[T any](ctx context.Context, c *redis.Client, key string) (T, error) {
	var dest T = *new(T)
	v, err := c.Get(ctx, key).Result()
	if err != nil {
		return dest, err
	}
	err = json.Unmarshal([]byte(v), &dest)
	if err != nil {
		return dest, err
	}
	return dest, nil
}
