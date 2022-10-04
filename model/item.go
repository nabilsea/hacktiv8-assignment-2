package model

type Item struct {
	ItemID      uint   `gorm:"primaryKey"`
	ItemCode    string `gorm:"type:varchar(50);"`
	Description string
	Quantity    uint
	OrderID     uint
}

func (Item) TableName() string {
	return "items"
}
