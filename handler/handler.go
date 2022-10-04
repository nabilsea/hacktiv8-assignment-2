package handler

import s "github.com/nabilsea/hacktiv8-assignment-2.git/service"

type Handler struct {
	orderService s.OrderService
}

type HandlerConfig struct {
	OrderService s.OrderService
}

func NewHandler(c *HandlerConfig) *Handler {
	return &Handler{
		orderService: c.OrderService,
	}
}
