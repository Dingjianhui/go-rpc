// Copyright © 2020 JianHui Ding
// go-grpc 场景练习
// github: https://github.com/Dingjianhui/go-rpc
// gitee:  https://gitee.com/dingjianhui/go-grpc

package services

import (
	"context"
	"fmt"
	. "grpc/pbfiles"
	)

type OrderService struct {

}
// 下订单
func (this *OrderService) CreateOrder(ctx context.Context, orderRequest *OrderRequest) (*OrderResponse, error)  {
	fmt.Println(orderRequest.OrderMain)

	// todo 订单创建

	// ......
	return &OrderResponse{
		Code: 1,
		Msg:  "订单创建成功",
	},nil
}
