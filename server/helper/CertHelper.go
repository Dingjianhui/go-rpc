package helper

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

// 获取服务端证书配置
func GetServerCreds() credentials.TransportCredentials  {
	//---------------自签证书-------------------//
	//creds,err := credentials.NewServerTLSFromFile("keys/server.crt","keys/server_no_passwd.key")
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//return creds
	//---------------自签证书-------------------//


	//---------------使用自签CA、server、Client证书和双向认证-------------------//
	// 创建证书池
	cert,err := tls.LoadX509KeyPair("cert/server.pem","cert/server.key")
	if err != nil {
		log.Fatal(err.Error())
	}
	certPool := x509.NewCertPool()
	ca,err := ioutil.ReadFile("cert/ca.pem")
	if err != nil {
		log.Fatal(err.Error())
	}
	certPool.AppendCertsFromPEM(ca)
	// 证书对象 credentials.TransportCredentials
	creds := credentials.NewTLS(&tls.Config{
		Certificates:                []tls.Certificate{cert}, //服务端证书
		ClientAuth:                  tls.RequireAndVerifyClientCert, // 双向验证
		ClientCAs:                   certPool,
	})
	return creds
	//---------------使用自签CA、server、Client证书和双向认证-------------------//
}


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
