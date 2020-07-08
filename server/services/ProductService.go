// Copyright © 2020 JianHui Ding
// go-grpc 场景练习
// github: https://github.com/Dingjianhui/go-rpc
// gitee:  https://gitee.com/dingjianhui/go-grpc

package services

import (
	"context"
	. "grpc/pbfiles"
)

type ProductService struct {

}

// 获取商品详情
func (this *ProductService) GetProductDetail(context.Context, *ProductDetailRequest) (*ProductsModel, error) {
	// todo 获取商品详情
	// ......

	return &ProductsModel{
		ProdId:    100,
		ProdName:  "测试商品",
		ProdDesc:  "商品描述",
		ProdPrice: 100.00,
	},nil
}
// 获取商品列表
func (this *ProductService) GetProductList(context.Context, *ProductListRequest) (*ProductListResponse, error) {
	// todo 获取商品列表
	// ......

	ret := []*ProductsModel {
		&ProductsModel{
			ProdId:    100,
			ProdName:  "测试商品",
			ProdDesc:  "商品描述",
			ProdPrice: 100.00,
		},
		&ProductsModel{
			ProdId:    101,
			ProdName:  "测试商品1",
			ProdDesc:  "商品描述1",
			ProdPrice: 101.00,
		},
	}
	return &ProductListResponse{ProdList:ret},nil
}
