// Copyright © 2020 JianHui Ding
// go-grpc 场景练习
// github: https://github.com/Dingjianhui/go-rpc
// gitee:  https://gitee.com/dingjianhui/go-grpc

package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"grpc/client/services"
	"grpc/client/utils"
	pb "grpc/pbfiles"
	"log"
	"time"
)

func main()  {

	creds := utils.GetClientCreds() // 获取客户端证书配置

	conn,err := grpc.Dial(":8081",grpc.WithTransportCredentials(creds)) // 使用证书连接服务端
	//conn,err := grpc.Dial(":8081",grpc.WithInsecure()) // 关闭Https安全传输
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	ctx := context.Background()

	//-------------商品管理--------------------//
	{
		productService := services.NewProducts(conn)

		// 获取商品列表
		req := &pb.ProductListRequest{
			Page:     1,
			PageSize: 10,
		}
		productService.GetProductList(ctx,req)

		// 获取商品详情
		dreq := &pb.ProductDetailRequest{ProdId:32}
		productService.GetProductDetail(ctx,dreq)
	}


	//-------------订单管理--------------------//
	// 测试数据
	//{"order_no":"20200703002","order_price":80,"user_id":3,"order_details":[{"order_no":"20200703002","prod_id":"101","prod_price":"20.0","buy_num":"3"}]}
	{
		orderService := services.NewOrders(conn)

		// 模拟请求数据
		t := timestamp.Timestamp{Seconds: time.Now().Unix()}
		requestBody := &pb.OrderRequest{OrderMain: &pb.OrdersModel{
			OrderNo:     "20200706001",
			UserId:      10,
			OrderPrice:  300.00,
			OrderStatus: 0,
			OrderTime:   &t,
			OrderDetails: []*pb.OrderDetailsModel{
				&pb.OrderDetailsModel{
					OrderNo:   "20200706001",
					ProdId:    20,
					ProdPrice: 300.00,
					BuyNum:    1,
				},
			},
		}}

		// 创建订单
		orderService.CreateOrder(ctx,requestBody)
	}





	//-------------批量获取用户积分--------------------//
	{
		userService := services.NewUsers(conn)

		// 客服端一次请求, 服务器一次应答
		userService.GetUserScore(ctx)

		// 服务端流
		//userService.GetUserScoreByServerStream(ctx)

		// 客户端流
		//userService.GetUserScoreByClientStream(ctx)

		// 双向流
		//userService.GetUserScoreByDoubleStream(ctx)
	}
}
