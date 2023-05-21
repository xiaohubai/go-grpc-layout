package middleware

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"

	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/internal/ecode"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

func Limiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		limiter := redis_rate.NewLimiter(consts.RDB)
		uri := c.Request.RequestURI
		index := strings.Index(uri, "?")
		var key string
		if index == -1 {
			key = uri
		} else {
			key = uri[:index]
		}
		res, err := limiter.Allow(c, key, redis_rate.PerMinute(int(consts.Conf.Limiter.Rate)))
		if err != nil {
			response.Fail(c, ecode.RateLimitAllowErrFailed, nil)
			c.Abort()
			return
		}
		c.Header("RateLimit-Remaining", strconv.Itoa(res.Remaining))
		if res.Allowed == 0 {
			seconds := int(res.RetryAfter / time.Second)
			c.Header("RateLimit-RetryAfter", strconv.Itoa(seconds))
			response.Fail(c, ecode.RateLimitAllowFailed, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
