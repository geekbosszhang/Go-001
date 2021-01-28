package data

import (
	"biz"
)
var _ biz.OrderRepo = (*orderRepo)(nil)

type orderRepo struct {}

func NewOrderRepo() biz.OrderRepo {
	return new(orderRepo)
}
