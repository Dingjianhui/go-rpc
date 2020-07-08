// Copyright © 2020 JianHui Ding
// go-grpc 场景练习
// github: https://github.com/Dingjianhui/go-rpc
// gitee:  https://gitee.com/dingjianhui/go-grpc

package services

import (
	"context"
	"fmt"
	pb "grpc/pbfiles"
	"io"
	"log"
	"time"
)

type UserService struct {

}

// 客服端一次请求, 服务器一次应答
func (this *UserService) GetUserScore(ctx context.Context, req *pb.UserScoreRequest) (*pb.UserScoreResponse, error)  {
	// todo 获取用户积分
	// ......


	// 模拟获取用户积分
	var score int32 = 101
	users := make([]*pb.UserInfo,0)
	for _,user := range req.Users {
		user.UserScore = score
		score++
		users = append(users,user)
	}

	return &pb.UserScoreResponse{Users:users},nil

}



// 客服端一次请求, 服务器多次应答(流式)
func (this *UserService) GetUserScoreByServerStream(req *pb.UserScoreRequest, stream pb.UserService_GetUserScoreByServerStreamServer) error {
	// todo 获取用户积分
	// ......

	// 模拟获取用户积分
	var score int32 = 101
	users := make([]*pb.UserInfo,0)
	for index,user := range req.Users {

		// 模拟查询通过用户ID查询用户积分
		time.Sleep(time.Second*2) // 模拟查询时用时很久
		user.UserScore = score
		score++
		users = append(users,user)

		// 模拟查询2条就发送
		if index > 0 && (index + 1) % 2 == 0 {
			err := stream.Send(&pb.UserScoreResponse{Users:users})
			if err != nil {
				return err
			}
			// 发送完后users重置
			users = (users)[0:0]
		}
	}

	// 处理余留未发送的
	if len(users) > 0 {
		err := stream.Send(&pb.UserScoreResponse{Users:users})
		if err != nil {
			return err
		}
	}

	return nil

}

// 客服端多次请求(流式), 服务器一次应答
func (this *UserService) GetUserScoreByClientStream(stream pb.UserService_GetUserScoreByClientStreamServer) error {
	// todo 获取用户积分
	// ......

	fmt.Println(123)
	// 模拟获取用户积分
	var score int32 = 101
	users := make([]*pb.UserInfo,0)

	for {
		userScoerReq,err := stream.Recv()

		if err == io.EOF { // 接收完了
			return stream.SendAndClose(&pb.UserScoreResponse{Users:users}) // 结束 后返回给客户端
		}

		if err != nil {
			return err
		}

		// 模拟查询用户积分
		if len(userScoerReq.Users) > 0 {
			for _,user := range userScoerReq.Users {
				user.UserScore = score
				score++
				users = append(users,user)
			}
		}
		fmt.Println("接收成功")
	}

}

// 客服端多次请求(流式), 服务器多次应答(流式)
func (this *UserService) GetUserScoreByDoubleStream(stream pb.UserService_GetUserScoreByDoubleStreamServer) error {
	// todo 获取用户积分
	// ......

	// 模拟获取用户积分
	var score int32 = 101
	users := make([]*pb.UserInfo,0)
	// 循环获取用户发送过来的信息
	for {
		req,err := stream.Recv()
		if err == io.EOF {
			// 接收完了
			return nil
		}
		if err != nil {
			log.Fatal(err.Error())
		}
		for _,user := range req.Users {
			time.Sleep(time.Second*1) //模拟查询积分耗时
			// 模拟查询积分
			user.UserScore = score
			score++
			users = append(users,user)
		}
		// 发送查询好的用户积分
		err = stream.Send(&pb.UserScoreResponse{Users:users})
		if err != nil {
			log.Println(err.Error())
		}
		// 将发送好的用户清空
		users = (users)[0:0]
	}

	return nil
}