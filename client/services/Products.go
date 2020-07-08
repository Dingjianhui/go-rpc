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

type Products struct {
	client pb.ProductServiceClient
}

func NewProducts(conn *grpc.ClientConn) *Products {
	cli := pb.NewProductServiceClient(conn)
	return &Products{client:cli}
}

// 商品列表
func (this *Products) GetProductList(ctx context.Context,req *pb.ProductListRequest)  {
	// 获取商品列表
	prodListRes,err := this.client.GetProductList(ctx,req)

	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(prodListRes)
}

// 商品详情
func (this *Products) GetProductDetail(ctx context.Context,req *pb.ProductDetailRequest)  {
	prodRes,err := this.client.GetProductDetail(ctx,req)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(prodRes)
}
