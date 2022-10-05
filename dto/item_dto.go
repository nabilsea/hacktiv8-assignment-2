package dto

import "github.com/nabilsea/hacktiv8-assignment-2.git/model"

type CreateItemRequest struct {
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity" binding:"required,min=0"`
	OrderID     uint
}

type UpdateItemRequest struct {
	LineItemID  uint   `json:"lineItemId" binding:"required"`
	ItemCode    string `json:"itemCode" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    uint   `json:"quantity" binding:"required,min=0"`
	OrderID     uint
}

type ItemResponse struct {
	LineItemID  uint   `json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"orderId"`
}

func FormatItem(item model.Item) ItemResponse {
	return ItemResponse{
		LineItemID:  item.ItemID,
		ItemCode:    item.ItemCode,
		Description: item.Description,
		Quantity:    item.Quantity,
		OrderID:     item.OrderID,
	}
}

func FormatItems(items []model.Item) []ItemResponse {
	formattedItems := []ItemResponse{}
	for _, item := range items {
		formattedItem := FormatItem(item)
		formattedItems = append(formattedItems, formattedItem)
	}
	return formattedItems
}
