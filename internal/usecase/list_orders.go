package usecase

import (
	"github.com/mersonff/desafio-03-go-expert/internal/entity"
	"github.com/mersonff/desafio-03-go-expert/pkg/events"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(orderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (l *ListOrdersUseCase) Execute() ([]*OrderOutputDTO, error) {
	orders, err := l.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var orderDTOs []*OrderOutputDTO
	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		}
		orderDTOs = append(orderDTOs, &dto)
	}

	return orderDTOs, nil
}

type ListOrdersUseCaseInterface interface {
	Execute() ([]*OrderOutputDTO, error)
}

type OrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}
