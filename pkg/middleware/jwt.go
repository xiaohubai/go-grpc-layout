package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/errors/v1"
	"github.com/xiaohubai/go-grpc-layout/pkg/jwt"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.Split(c.Request.Header.Get("Authorization"), " ")
		if len(auth) != 2 {
			response.Fail(c, v1.Error_TokenFailed, "缺少token", nil)
			c.Abort()
			return
		}
		token := auth[1]
		claims, err := jwt.Parse(token)
		if err != nil {
			response.Fail(c, v1.Error_TokenFailed, err)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
