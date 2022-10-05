package dto

type CreateItemRequest struct {
	ItemCode    string `json:"itemCode" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    uint   `json:"quantity" binding:"required,min=0"`
}

type UpdateItemRequest struct {
	LineItemID  uint   `json:"lineItemId" binding:"required"`
	ItemCode    string `json:"itemCode" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    uint   `json:"quantity" binding:"required,min=0"`
}
