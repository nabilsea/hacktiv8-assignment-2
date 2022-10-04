package model

import "time"

type Order struct {
	OrderID      uint   `gorm:"primaryKey"`
	CustomerName string `gorm:"type:varchar(50);"`
	OrderedAt    time.Time
	Item         []Item
}

func (Order) TableName() string {
	return "orders"
}
