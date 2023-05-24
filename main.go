package main

import (
	"github.com/xiaohubai/go-grpc-layout/cmd"
)

// @title Swagger Example API
// @version 0.0.1
// @description 总入口
// @in header
// @BasePath /
func main() {
	app := cmd.NewApp()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
