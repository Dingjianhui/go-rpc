package main

import (
	"google.golang.org/grpc"
	"grpc/server/helper"
	"grpc/server/services"
	"log"
	"net"
)

func main()  {

	creds := helper.GetServerCreds() // 获取服务端证书配置

	rpcServer := grpc.NewServer(grpc.Creds(creds))

	services.RegisterProdServiceServer(rpcServer,new(services.ProdService))

	lis,err := net.Listen("tcp",":8081")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = rpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err.Error())
	}

}
