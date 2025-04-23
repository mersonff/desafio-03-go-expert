package grpc

import (
	"context"

	"github.com/mersonff/desafio-03-go-expert/api/proto"
	"github.com/mersonff/desafio-03-go-expert/internal/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderServer struct {
	proto.UnimplementedOrderServiceServer
	orderUseCase domain.OrderUseCase
}

func NewOrderServer(orderUseCase domain.OrderUseCase) *OrderServer {
	return &OrderServer{
		orderUseCase: orderUseCase,
	}
}

func (s *OrderServer) ListOrders(ctx context.Context, req *proto.ListOrdersRequest) (*proto.ListOrdersResponse, error) {
	orders, err := s.orderUseCase.ListOrders()
	if err != nil {
		return nil, err
	}

	var protoOrders []*proto.Order
	for _, order := range orders {
		protoOrders = append(protoOrders, &proto.Order{
			Id:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
			CreatedAt:  timestamppb.New(order.CreatedAt).String(),
			UpdatedAt:  timestamppb.New(order.UpdatedAt).String(),
		})
	}

	return &proto.ListOrdersResponse{
		Orders: protoOrders,
	}, nil
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	order, err := s.orderUseCase.CreateOrder(req.Price, req.Tax)
	if err != nil {
		return nil, err
	}

	return &proto.CreateOrderResponse{
		Order: &proto.Order{
			Id:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
			CreatedAt:  timestamppb.New(order.CreatedAt).String(),
			UpdatedAt:  timestamppb.New(order.UpdatedAt).String(),
		},
	}, nil
}
