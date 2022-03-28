package order

type orderModel struct {
	OrderID int    `json:"order_id" validate:"required"`
	Price   int    `json:"price" validate:"required"`
	Title   string `json:"title" validate:"required"`
}
