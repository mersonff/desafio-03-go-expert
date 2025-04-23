package domain

import (
	"time"
)

type Order struct {
	ID          string
	Price       float64
	Tax         float64
	FinalPrice  float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrderRepository interface {
	Save(order *Order) error
	List() ([]*Order, error)
}

type OrderUseCase interface {
	CreateOrder(price, tax float64) (*Order, error)
	ListOrders() ([]*Order, error)
} 