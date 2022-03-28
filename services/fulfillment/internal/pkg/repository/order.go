package repository

import (
	"encoding/json"
	"github.com/adjust/rmq/v4"
)

type Order struct {
	OrderID int    `json:"order_id"`
	Price   int    `json:"price"`
	Title   string `json:"title"`
}

type orderRepository struct {
	queue rmq.Queue
}

func (r orderRepository) Create(order *Order) error {
	bytes, err := json.Marshal(order)
	if err != nil {
		return err
	}
	if err = r.queue.Publish(string(bytes)); err != nil {
		return err
	}
	return nil
}
