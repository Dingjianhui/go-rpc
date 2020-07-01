package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/client/helper"
	"grpc/client/services"
	"log"
)

func main()  {

	creds := helper.GetClientCreds() // 获取客户端证书配置

	conn,err := grpc.Dial(":8081",grpc.WithTransportCredentials(creds))
	//conn,err := grpc.Dial(":8081",grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	prodServiceClient := services.NewProdServiceClient(conn)
	prodRes,err := prodServiceClient.GetProdStock(context.Background(),&services.ProdRequest{ProdId:32})
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(prodRes.ProdStock)
}
