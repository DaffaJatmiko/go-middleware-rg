package main

import (
	"middleware/handler"
	"middleware/middleware"
	"middleware/repository"
	"middleware/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// wiring contract
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	handler := handler.NewHandler(svc)

	// routes
	v1 := r.Group("/v1", middleware.Auth(), middleware.Auth2())
	{
		v1.POST("/data", handler.CreateData)
		v1.GET("/data", handler.GetData)
		// v1.GET("/data/:id")
		// v1.PUT("/data/:id")
		// v1.DELETE("/data/:id")
	}

	r.Run() // 8080
}
