package main

import (
	"github.com/xiaohubai/go-grpc-layout/cmd"
)

func main() {
	app := cmd.NewApp()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
