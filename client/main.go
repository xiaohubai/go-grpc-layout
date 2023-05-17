package main

import (
	"context"
	"fmt"

	v1 "github.com/xiaohubai/go-grpc-layout/api/grpc/v1"
	conf "github.com/xiaohubai/go-grpc-layout/pkg/configs"
	"github.com/xiaohubai/go-grpc-layout/pkg/consul"
)

func main() {
	cc, err := conf.Load()
	if err != nil {
		panic("load config failed")
	}
	conn, err := consul.NewDiscovery(cc.Consul)
	if err != nil {
		panic(err)
	}
	client := v1.NewGrpcClient(conn)
	resp, err := client.GetUserInfo(context.Background(), &v1.UserInfoRequest{UserName: "s2ddds"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", resp)
}
