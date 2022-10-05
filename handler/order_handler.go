package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nabilsea/hacktiv8-assignment-2.git/dto"
	"github.com/nabilsea/hacktiv8-assignment-2.git/helper/util"
)

func (h *Handler) GetOrders(c *gin.Context) {
	orders, err := h.orderService.GetOrders()
	if err != nil {
		code := http.StatusBadRequest
		response := util.APIResponse(code, http.StatusText(code), err.Error(), nil)
		c.JSON(code, response)
		return
	}

	code := http.StatusOK
	formattedOrder := dto.FormatOrders(orders)
	response := util.APIResponse(code, http.StatusText(code), nil, formattedOrder)
	c.JSON(code, response)
}

func (h *Handler) CreateOrder(c *gin.Context) {
	input := &dto.CreateOrderRequest{}
	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := util.FormatValidationError(err)
		code := http.StatusUnprocessableEntity
		response := util.APIResponse(code, http.StatusText(code), errors, nil)
		c.JSON(code, response)
		return
	}

	order, err := h.orderService.CreateOrder(input)
	if err != nil {
		code := http.StatusBadRequest
		response := util.APIResponse(code, http.StatusText(code), err.Error(), nil)
		c.JSON(code, response)
		return
	}

	code := http.StatusOK
	formattedOrder := dto.FormatOrder(order)
	response := util.APIResponse(code, http.StatusText(code), nil, formattedOrder)
	c.JSON(code, response)
}

func (h *Handler) UpdateOrder(c *gin.Context) {

}

func (h *Handler) DeleteOrder(c *gin.Context) {

}
