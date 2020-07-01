package services

import "context"

type ProdService struct {

}

func (this *ProdService) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error)  {
	// todo 获取库存业务

	// ......

	stock := req.ProdId // 模拟演示，库存=商品id

	return &ProdResponse{ProdStock:stock}, nil
}
