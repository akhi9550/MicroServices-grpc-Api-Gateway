package main

import (
	"MicroServices-grpc/pkg/auth"
	"MicroServices-grpc/pkg/config"
	"MicroServices-grpc/pkg/order"
	"MicroServices-grpc/pkg/product"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
