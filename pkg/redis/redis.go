package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/xiaohubai/go-grpc-layout/configs/conf"
)

type Redis struct {
	Client *redis.Client
}

func NewClient(c *conf.Data_Redis) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       int(c.Db),
	})
	_, err := client.Ping(context.Background()).Result()
	return &Redis{Client: client}, err
}

func New(client *redis.Client) *Redis {
	return &Redis{Client: client}
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *Redis) Set(ctx context.Context, key string, value string) error {
	return r.Client.Set(ctx, key, value, 0).Err()
}

func (r *Redis) SetEx(ctx context.Context, key, value string, expiration time.Duration) error {
	return r.Client.SetEx(ctx, key, value, expiration).Err()
}

func (r *Redis) Lock(ctx context.Context, key, value string, expiration time.Duration) bool {
	ok, err := r.Client.SetNX(ctx, key, value, expiration).Result()
	if err != nil || !ok {
		return false
	}
	return true
}

func (r *Redis) UnLock(ctx context.Context, key, value string) bool {
	luaScript := `
        if redis.call("GET",KEYS[1]) == ARGV[1] then
            return redis.call("DEL",KEYS[1])
        else
            return 0
        end
	`
	ret, err := r.Client.Eval(ctx, luaScript, []string{key}, value).Result()
	if err != nil {
		return false
	}
	if v, ok := ret.(int64); ok {
		return v == 1
	}
	return false
}
