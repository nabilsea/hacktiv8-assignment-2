package service

import (
	"github.com/nabilsea/hacktiv8-assignment-2.git/model"
	r "github.com/nabilsea/hacktiv8-assignment-2.git/repository"
)

type OrderService interface {
	GetOrders() ([]*model.Order, error)
	GetOrderById() (*model.Order, error)
	CreateOrder() (*model.Order, error)
	UpdateOrder() (*model.Order, error)
	DeleteOrder() (*model.Order, error)
}

type orderService struct {
	orderRepository r.OrderRepository
}

type OSConfig struct {
	OrderRepository r.OrderRepository
}

func NewOrderService(c *OSConfig) OrderService {
	return &orderService{
		orderRepository: c.OrderRepository,
	}
}

func (s *orderService) GetOrders() ([]*model.Order, error) {
	return nil, nil
}

func (s *orderService) GetOrderById() (*model.Order, error) {
	return nil, nil
}

func (s *orderService) CreateOrder() (*model.Order, error) {
	return nil, nil
}

func (s *orderService) UpdateOrder() (*model.Order, error) {
	return nil, nil
}

func (s *orderService) DeleteOrder() (*model.Order, error) {
	return nil, nil
}
