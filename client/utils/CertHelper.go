// Copyright © 2020 JianHui Ding
// go-grpc 场景练习
// github: https://github.com/Dingjianhui/go-rpc
// gitee:  https://gitee.com/dingjianhui/go-grpc

package utils

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

// 获取客户端证书配置
func GetClientCreds() credentials.TransportCredentials {
	//---------------自签证书-------------------//
	//creds,err := credentials.NewClientTLSFromFile("keys/server.crt","ttphp.cn")
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//return creds
	//---------------自签证书-------------------//


	//---------------使用自签CA、server、Client证书和双向认证-------------------//
	cert,err := tls.LoadX509KeyPair("cert/client.pem","cert/client.key")
	if err != nil {
		log.Fatal(err.Error())
	}
	certPool := x509.NewCertPool()
	ca,err := ioutil.ReadFile("cert/ca.pem")
	if err != nil {
		log.Fatal(err.Error())
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates:                []tls.Certificate{cert}, // 加载客户端证书
		ServerName:                  "localhost",
		RootCAs:                     certPool,
	})
	return creds
	//---------------使用自签CA、server、Client证书和双向认证-------------------//

}