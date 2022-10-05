package dto

import "time"

type CreateOrderRequest struct {
	OrderedAt    time.Time           `json:"orderedAt" binding:"required"`
	CustomerName string              `json:"customerName" binding:"required"`
	Items        []CreateItemRequest `json:"items" binding:"required"`
}

type UpdateOrderRequest struct {
	OrderedAt    time.Time           `json:"orderedAt" binding:"required"`
	CustomerName string              `json:"customerName" binding:"required"`
	Items        []UpdateItemRequest `json:"items" binding:"required"`
}
