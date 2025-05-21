package usecase

import (
	"context"
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type ListOrdersInputDTO struct {
	Page  int
	Limit int
}

type ListOrdersOutputDTO struct {
	Orders []OrderOutputDTO
	Total  int
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

type ListOrdersUseCaseInterface interface {
	Execute(ctx context.Context, input ListOrdersInputDTO) (*ListOrdersOutputDTO, error)
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (uc *ListOrdersUseCase) Execute(ctx context.Context, input ListOrdersInputDTO) (*ListOrdersOutputDTO, error) {
	orders, err := uc.OrderRepository.List(ctx, input.Page, input.Limit)
	if err != nil {
		return nil, err
	}

	var output ListOrdersOutputDTO
	output.Orders = make([]OrderOutputDTO, len(orders))
	for i, order := range orders {
		output.Orders[i] = OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
	}

	total, err := uc.OrderRepository.Count(ctx)
	if err != nil {
		return nil, err
	}
	output.Total = total

	return &output, nil
} 