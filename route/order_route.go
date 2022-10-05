package route

import (
	"github.com/gin-gonic/gin"
	"github.com/nabilsea/hacktiv8-assignment-2.git/handler"
)

func (r *Router) Order(route *gin.RouterGroup, h *handler.Handler) {
	route = route.Group("/orders")
	route.GET("/", h.GetOrders)
	route.POST("/", h.CreateOrder)
	route.PUT("/:orderId", h.UpdateOrder)
	route.DELETE("/:orderId", h.DeleteOrder)
}
