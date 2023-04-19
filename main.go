package main

import "github.com/xiaohubai/go-grpc-layout/cmd"

func main() {
	app, close := cmd.Run()
	defer close()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
