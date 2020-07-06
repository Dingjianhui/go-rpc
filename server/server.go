package main

import (
	"google.golang.org/grpc"
	pb "grpc/pbfiles"
	"grpc/server/helper"
	"grpc/server/services"
	"log"
	"net"
)

func main()  {

	creds := helper.GetServerCreds() // 获取服务端证书配置

	rpcServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterProductServiceServer(rpcServer,new(services.ProductService)) // 注册产品服务

	pb.RegisterOrderServiceServer(rpcServer,new(services.OrderService)) // 注册订单服务

	pb.RegisterUserServiceServer(rpcServer,new(services.UserService)) // 注册用户服务

	lis,err := net.Listen("tcp",":8081")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = rpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err.Error())
	}

}
