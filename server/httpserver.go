package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"grpc/server/helper"
	"grpc/server/services"
	"log"
	"net/http"
)

func main()  {
	mux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCreds())}
	err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(),
		mux,
		"localhost:8081",
		opt,
		)

	if err != nil {
		log.Fatal(err.Error())
	}

	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
