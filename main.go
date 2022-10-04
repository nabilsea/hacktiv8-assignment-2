package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nabilsea/hacktiv8-assignment-2.git/config"
	"github.com/nabilsea/hacktiv8-assignment-2.git/handler"
	"github.com/nabilsea/hacktiv8-assignment-2.git/repository"
	"github.com/nabilsea/hacktiv8-assignment-2.git/route"
	"github.com/nabilsea/hacktiv8-assignment-2.git/service"
)

func main() {
	db := config.GetConn()

	orderRepository := repository.NewOrderRepository(&repository.ORConfig{DB: db})
	orderService := service.NewOrderService(&service.OSConfig{OrderRepository: orderRepository})
	h := handler.NewHandler(&handler.HandlerConfig{
		OrderService: orderService,
	})

	routes := route.NewRouter(&route.RouterConfig{})
	router := gin.Default()
	router.NoRoute(h.NoRoute)
	version := os.Getenv("API_VERSION")
	api := router.Group(fmt.Sprintf("/api/%s", version))
	routes.Order(api, h)

	router.Run(":8000")
}
