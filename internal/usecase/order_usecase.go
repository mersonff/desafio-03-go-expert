package usecase

import (
	"github.com/google/uuid"
	"github.com/mersonff/20-CleanArch/internal/domain"
	"time"
)

type orderUseCase struct {
	orderRepository domain.OrderRepository
}

func NewOrderUseCase(repo domain.OrderRepository) domain.OrderUseCase {
	return &orderUseCase{
		orderRepository: repo,
	}
}

func (u *orderUseCase) CreateOrder(price, tax float64) (*domain.Order, error) {
	order := &domain.Order{
		ID:         uuid.New().String(),
		Price:      price,
		Tax:        tax,
		FinalPrice: price + tax,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := u.orderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (u *orderUseCase) ListOrders() ([]*domain.Order, error) {
	return u.orderRepository.List()
} 