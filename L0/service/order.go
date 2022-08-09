package service

import (
	"L0/data/model"
	"L0/data/pg"
	"L0/data/redis"
	"encoding/json"
	"fmt"
)

type OrderService struct {
}

func NewOrderService() *OrderService {

	return &OrderService{}
}

func (o *OrderService) Set(order model.Order) {

	pg.AddOrder(order)
	orderB, _ := json.Marshal(order)
	err := redis.Set(order.OrderUid, string(orderB))
	if err != nil {
		fmt.Printf("error while set to Redis %s \n", err)
	}

}
