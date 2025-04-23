package graphql

import (
	"context"

	"github.com/mersonff/desafio-03-go-expert/internal/domain"
)

type Resolver struct {
	orderUseCase domain.OrderUseCase
}

func NewResolver(orderUseCase domain.OrderUseCase) *Resolver {
	return &Resolver{
		orderUseCase: orderUseCase,
	}
}

func (r *Resolver) Query() QueryResolver {
	return r
}

func (r *Resolver) Mutation() MutationResolver {
	return r
}

func (r *Resolver) Orders() ([]*Order, error) {
	orders, err := r.orderUseCase.ListOrders()
	if err != nil {
		return nil, err
	}

	var result []*Order
	for _, order := range orders {
		result = append(result, &Order{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
			CreatedAt:  order.CreatedAt.String(),
			UpdatedAt:  order.UpdatedAt.String(),
		})
	}

	return result, nil
}

func (r *Resolver) CreateOrder(ctx context.Context, input CreateOrderInput) (*Order, error) {
	order, err := r.orderUseCase.CreateOrder(input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	return &Order{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
		CreatedAt:  order.CreatedAt.String(),
		UpdatedAt:  order.UpdatedAt.String(),
	}, nil
}
