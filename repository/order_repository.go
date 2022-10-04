package repository

import (
	"github.com/nabilsea/hacktiv8-assignment-2.git/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	FindAll() ([]*model.Order, error)
	FindById(id int) (*model.Order, error)
	Save(order *model.Order) (*model.Order, error)
	Update(order *model.Order) (*model.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

type ORConfig struct {
	DB *gorm.DB
}

func NewOrderRepository(c *ORConfig) OrderRepository {
	return &orderRepository{
		db: c.DB,
	}
}

func (r *orderRepository) FindAll() ([]*model.Order, error) {
	var orders []*model.Order

	err := r.db.Find(&orders).Error
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (r *orderRepository) FindById(id int) (*model.Order, error) {
	var order *model.Order

	err := r.db.Where("id = ?", id).Find(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *orderRepository) Save(order *model.Order) (*model.Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *orderRepository) Update(order *model.Order) (*model.Order, error) {
	err := r.db.Save(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}
