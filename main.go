package main

import (
	"github.com/rama4zis/go-restapi-assignment2/models"

	"github.com/rama4zis/go-restapi-assignment2/controllers/productcontroller"

	"github.com/rama4zis/go-restapi-assignment2/controllers/ordercontroller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:id", productcontroller.Show)
	r.POST("/api/products", productcontroller.Create)
	r.PUT("/api/products/:id", productcontroller.Update)
	r.DELETE("/api/products", productcontroller.Delete)

	// for orders
	r.GET("/orders", ordercontroller.Index)
	// r.GET("/orders/:id", ordercontroller.Show)
	r.POST("/orders", ordercontroller.Create)
	r.PUT("/orders/:orderId", ordercontroller.Update)
	r.DELETE("/orders/:orderId", ordercontroller.Delete)

	r.Run()

}
