package main

import (
	"log"
	"net/http"

	"git.com/api-rest/handlers"
	"git.com/api-rest/repositories"
	"git.com/api-rest/services"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := repositories.NewMemoryRepository()
	service := services.NewRewardService(repo)
	handler := handlers.NewRewardHandler(service)

	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.POST("/purchases", handler.RegisterPurchase)
	v1.GET("/customers/:id/points", handler.GetPointsBalance)
	v1.POST("/customers/:id/redeem", handler.RedeemPoints)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor Gin: %v", err)
	}
}
