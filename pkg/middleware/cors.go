package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-grpc-layout/configs/conf"
)

func Cors(cors *conf.Cors) gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if cors.Mode == "allow-all" || origin == "" || origin == "null" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin,Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")

			if c.Request.Method == http.MethodOptions {
				c.AbortWithStatus(http.StatusNoContent)
			}
			c.Next()
		} else {
			whitelist := checkOrigin(origin, cors.Whitelist)
			if whitelist != nil {
				c.Header("Access-Control-Allow-Origin", whitelist.AllowOrigin)
				c.Header("Access-Control-Allow-Headers", whitelist.AllowHeaders)
				c.Header("Access-Control-Allow-Methods", whitelist.AllowMethods)
				c.Header("Access-Control-Expose-Headers", whitelist.ExposeHeaders)
				if whitelist.AllowCredentials {
					c.Header("Access-Control-Allow-Credentials", "true")
				}
			}
			if whitelist == nil {
				c.AbortWithStatus(http.StatusForbidden)
			}
			c.Next()
		}
	}
}

func checkOrigin(origin string, whitelist []*conf.Cors_Whitelist) *conf.Cors_Whitelist {
	for _, v := range whitelist {
		if origin == v.AllowOrigin {
			return v
		}
	}
	return nil
}
