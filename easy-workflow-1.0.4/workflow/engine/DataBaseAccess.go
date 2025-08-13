package engine

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// DB is the global redis client instance.
var DB *redis.Client

// Storage defines the redis operations used by the engine.
type Storage interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd
	Do(ctx context.Context, args ...interface{}) *redis.Cmd
}

// DBConnect initialises the redis client using options from DBConnConfigurator.
func DBConnect() error {
	DB = redis.NewClient(&redis.Options{
		Addr:         DBConnConfigurator.Addr,
		Password:     DBConnConfigurator.Password,
		DB:           DBConnConfigurator.DB,
		PoolSize:     DBConnConfigurator.PoolSize,
		MinIdleConns: DBConnConfigurator.MinIdleConns,
	})
	return DB.Ping(context.Background()).Err()
}

// RunCommand executes a redis command. When result is nil the raw result is returned.
// If result is non-nil, the command result will be scanned into it.
func RunCommand(cmd string, result interface{}, params ...interface{}) (interface{}, error) {
	args := append([]interface{}{cmd}, params...)
	res := DB.Do(context.Background(), args...)
	if result == nil {
		return res.Result()
	}
	if err := res.Scan(result); err != nil {
		return nil, err
	}
	return result, nil
}

// Convenience wrappers for common operations.
func Get(key string) (string, error) {
	return DB.Get(context.Background(), key).Result()
}

func Set(key string, value interface{}, expiration time.Duration) error {
	return DB.Set(context.Background(), key, value, expiration).Err()
}

func HGetAll(key string) (map[string]string, error) {
	return DB.HGetAll(context.Background(), key).Result()
}
