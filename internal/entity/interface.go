package entity

import "context"

type OrderRepositoryInterface interface {
	Save(order *Order) error
	List(ctx context.Context, page, limit int) ([]Order, error)
	Count(ctx context.Context) (int, error)
}
