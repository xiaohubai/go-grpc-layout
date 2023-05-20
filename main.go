package main

import (
	"github.com/xiaohubai/go-grpc-layout/cmd"
)

func main() {
	if err := cmd.NewApp().Run(); err != nil {
		panic(err)
	}
}
