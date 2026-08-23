package handlers

import (
	"net/http"

	"git.com/api-rest/models"
	"git.com/api-rest/services"
	"github.com/gin-gonic/gin"
)

type RewardHandler struct {
	service services.RewardService
}

func NewRewardHandler(service services.RewardService) *RewardHandler {
	return &RewardHandler{service: service}
}

func (h *RewardHandler) RegisterPurchase(c *gin.Context) {
	var req models.PurchaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid JSON Request Body"})
		return
	}

	if req.CustomerID == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "The customer_id field is required"})
		return
	}

	response, err := h.service.RegisterPurchase(req)
	if err != nil {
		switch err {
		case models.ErrInvalidAmount:
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h *RewardHandler) GetPointsBalance(c *gin.Context) {
	customerID := c.Param("id")
	if customerID == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "The customer ID is required"})
		return
	}

	response, err := h.service.GetCustomerPoints(customerID)
	if err != nil {
		switch err {
		case models.ErrCustomerNotFound:
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *RewardHandler) RedeemPoints(c *gin.Context) {
	customerID := c.Param("id")
	if customerID == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "The customer ID is required"})
		return
	}

	var req models.RedeemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid JSON Request Body"})
		return
	}

	response, err := h.service.RedeemPoints(customerID, req)
	if err != nil {
		switch err {
		case models.ErrCustomerNotFound:
			c.JSON(http.StatusNotFound, models.ErrorResponse{Error: err.Error()})
		case models.ErrInsufficientPoints, models.ErrInvalidPoints:
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, response)
}
