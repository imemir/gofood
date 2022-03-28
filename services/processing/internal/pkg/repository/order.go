package repository

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderID int    `json:"order_id"`
	Price   int    `json:"price"`
	Title   string `json:"title"`
}

type orderRepository struct {
	db *gorm.DB
}

func (r orderRepository) Save(order *Order) error {
	if err := r.db.Save(order).Error; err != nil {
		return err
	}
	return nil
}
