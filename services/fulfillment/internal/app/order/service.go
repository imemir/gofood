package order

import "github.com/imemir/gofood/services/fulfillment/internal/pkg/repository"

var Service = service{}

type service struct {
}

func (s service) Create(orderID int, price int, title string) error {
	if err := repository.Orders.Create(&repository.Order{
		OrderID: orderID,
		Price:   price,
		Title:   title,
	}); err != nil {
		return err
	}
	return nil
}
