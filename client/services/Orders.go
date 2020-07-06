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


