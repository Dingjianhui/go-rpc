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
	"io"
	"log"
)

type Users struct {
	client pb.UserServiceClient
}

func NewUsers(conn *grpc.ClientConn) *Users {
	cli := pb.NewUserServiceClient(conn)
	return &Users{client:cli}
}


//-------------批量获取用户积分--------------------//
// 客服端一次请求, 服务器一次应答
func (this *Users) GetUserScore(ctx context.Context)  {

	// 模拟批量用户ID
	req := &pb.UserScoreRequest{}
	var i int32
	req.Users = make([]*pb.UserInfo,0)
	for i = 1; i <= 5; i++ {
		req.Users = append(req.Users,&pb.UserInfo{UserId:i})
	}

	userRes,err := this.client.GetUserScore(ctx,req) // 批量发送
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("用户积分:")
	fmt.Println(userRes)
}

// 服务端流-----客服端一次请求, 服务器多次应答(流式)
func (this *Users) GetUserScoreByServerStream(ctx context.Context) {

	// 模拟批量用户ID
	req := &pb.UserScoreRequest{}
	var i int32
	req.Users = make([]*pb.UserInfo,0)
	for i = 1; i <= 5; i++ {
		req.Users = append(req.Users,&pb.UserInfo{UserId:i})
	}

	stream,err := this.client.GetUserScoreByServerStream(ctx,req) // 批量发送给服务端
	if err != nil {
		log.Fatal(err.Error())
	}

	// 循环获取服务端分批返回的信息
	for {
		userRes,err := stream.Recv()
		if err != nil {
			log.Fatal(err.Error())
		}
		if err == io.EOF {
			break
		}

		//  获取到部分数据后，可通过协程进行相关操作
		if len(userRes.Users) > 0 {
			for _,user := range userRes.Users {
				fmt.Printf("用户ID %d, 用户积分 %d \n ",user.UserId,user.UserScore)
			}
		}

	}
}

// 客户端流-----客服端多次请求(流式), 服务器一次应答
func (this *Users) GetUserScoreByClientStream(ctx context.Context) {

	stream,err := this.client.GetUserScoreByClientStream(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 模拟批量用户ID
	req := &pb.UserScoreRequest{}
	var i int32


	// 假设耗时，那么分批发送给服务端
	for j := 1; j <= 3; j++ {
		req.Users = make([]*pb.UserInfo,0)
		for i = 1; i <= 5; i++ {
			req.Users = append(req.Users,&pb.UserInfo{UserId:i})
		}
		err = stream.Send(req)
		if err != nil {
			log.Println(err)
		}
	}

	// 接收服务端返回信息
	userRes,err := stream.CloseAndRecv()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(userRes.Users)

}


// 双向流-----客服端多次请求(流式), 服务器多次应答(流式)
func (this *Users) GetUserScoreByDoubleStream(ctx context.Context)  {

	stream,err := this.client.GetUserScoreByDoubleStream(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 模拟批量用户ID
	req := &pb.UserScoreRequest{}
	var i int32


	// 假设耗时，那么分批发送给服务端
	for j := 1; j <= 3; j++ {
		req.Users = make([]*pb.UserInfo,0)
		for i = 1; i <= 2; i++ {
			req.Users = append(req.Users,&pb.UserInfo{UserId:i})
		}
		// 发送用户信息给服务端
		err = stream.Send(req)
		if err != nil {
			log.Println(err)
		}

		userScoreReq,err := stream.Recv()
		if err == io.EOF {
			// 接收完成
			break
		}
		if err != nil {
			log.Println(err)
		}

		// 打印接收到的信息
		fmt.Println(userScoreReq.Users)
	}
}





