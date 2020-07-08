// Copyright © 2020 JianHui Ding
// go-grpc 场景练习
// github: https://github.com/Dingjianhui/go-rpc
// gitee:  https://gitee.com/dingjianhui/go-grpc

package services

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc/pbfiles"
	"log"
)

type Orders struct {
	client pb.OrderServiceClient
}

func NewOrders(conn *grpc.ClientConn) *Orders {
	cli := pb.NewOrderServiceClient(conn)
	return &Orders{client:cli}
}

// 创建订单
func (this *Orders) CreateOrder(ctx context.Context,requestBody *pb.OrderRequest) {
	orderRes, err := this.client.CreateOrder(ctx, requestBody)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(orderRes)
}


