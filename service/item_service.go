package service

import (
	"github.com/nabilsea/hacktiv8-assignment-2.git/dto"
	"github.com/nabilsea/hacktiv8-assignment-2.git/helper/custom_error"
	"github.com/nabilsea/hacktiv8-assignment-2.git/model"
	r "github.com/nabilsea/hacktiv8-assignment-2.git/repository"
)

type ItemService interface {
	CreateItem(input *dto.CreateItemRequest) (*model.Item, error)
	UpdateItem() (*model.Item, error)
}

type itemService struct {
	itemRepository r.ItemRepository
}

type ISConfig struct {
	ItemRepository r.ItemRepository
}

func NewItemService(c *ISConfig) ItemService {
	return &itemService{
		itemRepository: c.ItemRepository,
	}
}

func (s *itemService) CreateItem(input *dto.CreateItemRequest) (*model.Item, error) {
	item, err := s.itemRepository.FindByCode(input.ItemCode)
	if err != nil {
		return item, err
	}
	if item.ItemID != 0 {
		return item, &custom_error.ItemAlreadyExists{}
	}

	item.ItemCode = input.ItemCode
	item.Description = input.Description
	item.Quantity = input.Quantity
	item.OrderID = input.OrderID

	item, err = s.itemRepository.Save(item)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (s *itemService) UpdateItem() (*model.Item, error) {
	return nil, nil
}
