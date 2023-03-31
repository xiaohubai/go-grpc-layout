package middleware

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
)

// Server is an server logging middleware.
func Logger(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {

			reply, err = handler(ctx, req)

			log.WithContext(ctx, logger).Log(log.LevelInfo,
				"kind", "server",
			)
			return
		}
	}
}
