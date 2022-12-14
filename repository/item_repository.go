package repository

import (
	"github.com/nabilsea/hacktiv8-assignment-2.git/model"
	"gorm.io/gorm"
)

type ItemRepository interface {
	FindAll() ([]*model.Item, error)
	FindById(id int) (*model.Item, error)
	FindByCode(code string) (*model.Item, error)
	Save(item *model.Item) (*model.Item, error)
	Update(item *model.Item) (*model.Item, error)
	Delete(item *model.Item) (*model.Item, error)
	DeleteByOrderId(orderID int) error
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

	err := r.db.Where("item_id = ?", id).Find(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *itemRepository) FindByCode(code string) (*model.Item, error) {
	var item *model.Item

	err := r.db.Where("item_code = ?", code).Find(&item).Error
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

func (r *itemRepository) Delete(item *model.Item) (*model.Item, error) {
	err := r.db.Delete(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *itemRepository) DeleteByOrderId(orderID int) error {
	err := r.db.Where("order_id = ?", orderID).Delete(&model.Item{}).Error
	if err != nil {
		return err
	}

	return nil
}
