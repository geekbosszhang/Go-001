package service

import (
	"context"
	"biz"
)

type ShopService struct {
	ouc *biz.OrderUsecase
}

func NewShopService(ouc *biz.OrderUsercase) ShopService {
	return &ShopService{
		ouc: ouc,
	}
}

func (svr *ShopService) CreateOrder(ctx context.Context) {

}
