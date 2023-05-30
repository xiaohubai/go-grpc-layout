package middleware

import (
	"sync"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"

	"github.com/xiaohubai/go-grpc-layout/configs/conf"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
	"github.com/xiaohubai/go-grpc-layout/internal/ecode"
	"github.com/xiaohubai/go-grpc-layout/pkg/jwt"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

func Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		userInfo := claims.(*jwt.Claims)
		obj := c.Request.URL.Path
		act := c.Request.Method
		sub := userInfo.RoleID
		e := SyncedEnforcer()
		if ok, err := e.Enforce(sub, obj, act); !ok {
			response.Fail(c, ecode.CasbinFailed, err)
			c.Abort()
			return
		}
		c.Next()
	}
}

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func SyncedEnforcer() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDBUseTableName(consts.DB, "", model.TableNameCasbinRule)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(conf.C.Casbin.Path, a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
