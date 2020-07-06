package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	pb "grpc/pbfiles"
	"grpc/server/helper"
	"log"
	"net/http"
)

func main()  {
	mux := runtime.NewServeMux() // 初始化路由

	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCreds())} // 使用证书连接

	//---------------注册产品服务--------------------//
	err := pb.RegisterProductServiceHandlerFromEndpoint(context.Background(),
		mux,
		"localhost:8081",
		opt,
		)

	if err != nil {
		log.Fatal(err.Error())
	}
	//---------------注册产品服务--------------------//



	//---------------注册订单服务--------------------//
	err = pb.RegisterOrderServiceHandlerFromEndpoint(context.Background(),
		mux,
		"localhost:8081",
		opt,
		)
	if err != nil {
		log.Fatal(err.Error())
	}
	//---------------注册订单服务--------------------//




	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
