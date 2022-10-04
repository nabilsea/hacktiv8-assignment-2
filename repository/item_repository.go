package repository

import (
	"github.com/nabilsea/hacktiv8-assignment-2.git/model"
	"gorm.io/gorm"
)

type ItemRepository interface {
	FindAll() ([]*model.Item, error)
	FindById(id int) (*model.Item, error)
	Save(item *model.Item) (*model.Item, error)
	Update(item *model.Item) (*model.Item, error)
}

type itemRepository struct {
	db *gorm.DB
}

type IRConfig struct {
	DB *gorm.DB
}

func NewItemRepository(c *IRConfig) ItemRepository {
	return &itemRepository{
		db: c.DB,
	}
}

func (r *itemRepository) FindAll() ([]*model.Item, error) {
	var items []*model.Item

	err := r.db.Find(&items).Error
	if err != nil {
		return items, err
	}

	return items, nil
}

func (r *itemRepository) FindById(id int) (*model.Item, error) {
	var item *model.Item

	err := r.db.Where("id = ?", id).Find(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *itemRepository) Save(item *model.Item) (*model.Item, error) {
	err := r.db.Create(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *itemRepository) Update(item *model.Item) (*model.Item, error) {
	err := r.db.Save(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil
}
