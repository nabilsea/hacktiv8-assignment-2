package dto

import "github.com/nabilsea/hacktiv8-assignment-2.git/model"

type CreateOrderRequest struct {
	OrderedAt    string              `json:"orderedAt" binding:"required"`
	CustomerName string              `json:"customerName" binding:"required"`
	Items        []CreateItemRequest `json:"items" binding:"required"`
}
type UpdateOrderRequest struct {
	OrderedAt    string              `json:"orderedAt" binding:"required"`
	CustomerName string              `json:"customerName" binding:"required"`
	Items        []UpdateItemRequest `json:"items" binding:"required"`
}

type OrderResponse struct {
	OrderID      uint           `json:"orderId"`
	OrderedAt    string         `json:"orderedAt"`
	CustomerName string         `json:"customerName"`
	Items        []ItemResponse `json:"items"`
}

func FormatOrder(order *model.Order) OrderResponse {
	return OrderResponse{
		OrderID:      order.OrderID,
		OrderedAt:    order.OrderedAt.String(),
		CustomerName: order.CustomerName,
		Items:        FormatItems(order.Items),
	}
}

func FormatOrders(orders []*model.Order) []OrderResponse {
	formattedOrders := []OrderResponse{}
	for _, order := range orders {
		formattedOrder := FormatOrder(order)
		formattedOrders = append(formattedOrders, formattedOrder)
	}
	return formattedOrders
}
