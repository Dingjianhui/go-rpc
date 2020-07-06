gRPC 场景练习 (Go)
======================

DESCRIPTION
-------------
go-grpc场景练习-模拟创建订单服务、产品列表服务、产品详情服务、使用流模式批量获取用户积分服务

BACKGROUND
-------------
生成服务端与客户端代码
[Product.proto](protos/Product.proto).
[Order.proto](protos/Order.proto).
[User.proto](protos/User.proto).

PREREQUISITES
-------------

- This requires Go 1.14 or later
- Requires that [GOPATH is set](https://golang.org/doc/code.html#GOPATH)

```
$ go help gopath
$ # ensure the PATH contains $GOPATH/bin
$ export PATH=$PATH:$GOPATH/bin
```

INSTALL
-------

```
git clone https://github.com/Dingjianhui/go-rpc.git
cd go-rpc
go mod tidy
```

TRY IT!
-------

- Run the server

  ```
  $ go run server/server.go
  ```

- Run the client

  ```
  $ go run client/client.go
  ```

OPTIONAL - Rebuilding the generated code
----------------------------------------

protoc安装文档 [protobuf compiler](https://github.com/google/protobuf/blob/master/README.md#protocol-compiler-installation)


1. 安装 Protocol buffer 编译器(protoc) V3版本 `https://github.com/protocolbuffers/protobuf/releases`
2. 安装 Protocol buffer 编译器的 Go插件 (protoc-gen-go) `go get github.com/golang/protobuf/protoc-gen-go`
3. 安装 Protocol buffer 编译器的插件-grpc-gateway `go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway`
4. 安装 Protocol buffer 编译器的插件-swagger  `go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger`


