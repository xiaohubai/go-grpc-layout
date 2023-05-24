package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisLock struct {
	rdb        *redis.Client
	key        string
	value      string
	expiration time.Duration
}

func NewRedisLock(rdb *redis.Client, key, value string, expiration time.Duration) *RedisLock {
	return &RedisLock{
		rdb:        rdb,
		key:        key,
		value:      value,
		expiration: expiration,
	}
}

func (l *RedisLock) Lock(ctx context.Context) bool {
	ok, err := l.rdb.SetNX(ctx, l.key, l.value, l.expiration).Result()
	if err != nil || !ok {
		return false
	}
	return true
}

func (l *RedisLock) UnLock(ctx context.Context, key, value string) bool {
	luaScript := `
        if redis.call("GET",KEYS[1]) == ARGV[1] then
            return redis.call("DEL",KEYS[1])
        else
            return 0
        end
	`
	ret, err := l.rdb.Eval(ctx, luaScript, []string{l.key}, l.value).Result()
	if err != nil {
		return false
	}
	if v, ok := ret.(int64); ok {
		return v == 1
	}
	return false
}
