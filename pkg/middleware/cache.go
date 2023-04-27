package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
)

func Cache() gin.HandlerFunc {
	return func(c *gin.Context) {
		/* 	if c.Request.Method != http.MethodGet {
			c.Next()
			return
		} */
		cacheKey := c.Request.URL.String()
		ctx := context.Background()
		cacheValue, err := consts.RDB.Get(ctx, cacheKey).Result()
		if err == nil {
			c.Header("Cache-Control", "max-age=3600")
			c.String(http.StatusOK, cacheValue)
			return
		}

		lockKey := cacheKey + ":lock"
		lockValue := fmt.Sprintf("%d", time.Now().UnixNano())
		lockAcquired, err := acquireLock(consts.RDB, lockKey, lockValue, 30)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		if !lockAcquired {
			// Failed to acquire lock
			c.Header("Cache-Control", "max-age=0") // Disable caching
			return
		}
		defer func() {
			err := releaseLock(consts.RDB, lockKey, lockValue)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
			}
		}()
		c.Next()
		if c.Writer.Status() == http.StatusOK {
			/* 	err := consts.RDB.Set(ctx, cacheKey, c.Writer.Body.String(), 1*time.Hour).Err()
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			} */
		}
	}
}

func acquireLock(rdb *redis.Client, lockKey string, lockValue string, lockTimeout int) (bool, error) {
	ctx := context.Background()
	result, err := rdb.SetNX(ctx, lockKey, lockValue, 0).Result()
	if err != nil {
		return false, err
	}
	if result == true {
		// Set lock timeout
		err = rdb.Expire(ctx, lockKey, time.Duration(lockTimeout)*time.Second).Err()
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

func releaseLock(rdb *redis.Client, lockKey string, lockValue string) error {
	ctx := context.Background()
	currentValue, err := rdb.Get(ctx, lockKey).Result()
	if err != nil {
		return err
	}
	if currentValue == lockValue {
		err = rdb.Del(ctx, lockKey).Err()
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("lock key is not owned by this process")
	}
	return nil
}
