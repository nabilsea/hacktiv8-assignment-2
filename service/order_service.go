package service

import (
	"log"
	"time"

	"github.com/nabilsea/hacktiv8-assignment-2.git/dto"
	"github.com/nabilsea/hacktiv8-assignment-2.git/helper/custom_error"
	"github.com/nabilsea/hacktiv8-assignment-2.git/model"
	r "github.com/nabilsea/hacktiv8-assignment-2.git/repository"
)

type OrderService interface {
	GetOrders() ([]*model.Order, error)
	GetOrderById() (*model.Order, error)
	CreateOrder(input *dto.CreateOrderRequest) (*model.Order, error)
	UpdateOrder() (*model.Order, error)
	DeleteOrder(params *dto.OrderParams) (*model.Order, error)
}

type orderService struct {
	orderRepository r.OrderRepository
	itemService     ItemService
}

type OSConfig struct {
	OrderRepository r.OrderRepository
	ItemService     ItemService
}

func NewOrderService(c *OSConfig) OrderService {
	return &orderService{
		orderRepository: c.OrderRepository,
		itemService:     c.ItemService,
	}
}

func (s *orderService) GetOrders() ([]*model.Order, error) {
	orders, err := s.orderRepository.FindAll()
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (s *orderService) GetOrderById() (*model.Order, error) {
	return nil, nil
}

func (s *orderService) CreateOrder(input *dto.CreateOrderRequest) (*model.Order, error) {
	// layoutFormat := "2006-01-02T15:04:05-0700"
	orderedAt, err := time.Parse(time.RFC3339, input.OrderedAt)
	if err != nil {
		return &model.Order{}, err
	}

	order := &model.Order{}
	order.CustomerName = input.CustomerName
	order.OrderedAt = orderedAt

	order, err = s.orderRepository.Save(order)
	if err != nil {
		return order, err
	}

	for _, inputItem := range input.Items {
		inputItem.OrderID = order.OrderID
		item, err := s.itemService.CreateItem(&inputItem)
		if err != nil {
			log.Println(err.Error())
		} else {
			order.Items = append(order.Items, *item)
		}
	}

	return order, nil
}

func (s *orderService) UpdateOrder() (*model.Order, error) {
	return nil, nil
}

func (s *orderService) DeleteOrder(params *dto.OrderParams) (*model.Order, error) {
	order, err := s.orderRepository.FindById(int(params.OrderID))
	if err != nil {
		return order, err
	}
	if order.OrderID == 0 {
		return order, &custom_error.OrderNotFound{}
	}

	err = s.itemService.DeleteItemByOrderId(order.OrderID)
	if err != nil {
		return order, err
	}

	order, err = s.orderRepository.Delete(order)
	if err != nil {
		return order, err
	}

	return order, nil
}
