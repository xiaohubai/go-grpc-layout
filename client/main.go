package main

import (
	"context"
	"fmt"

	v1 "github.com/xiaohubai/go-grpc-layout/api/grpc/v1"
	"github.com/xiaohubai/go-grpc-layout/pkg/consul"
	"github.com/xiaohubai/go-grpc-layout/pkg/viper"
)

func main() {
	cc, err := viper.Load()
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
