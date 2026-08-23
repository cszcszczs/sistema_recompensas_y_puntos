package models

import (
	"errors"
	"time"
)

var (
	ErrInsufficientPoints = errors.New("Insuficient balance of points to redeem")
	ErrCustomerNotFound   = errors.New("Customer not found")
	ErrInvalidAmount      = errors.New("The amount entered must be greater than 0")
	ErrInvalidPoints      = errors.New("The number of points entered must be greater than 0")
)

const (
	PesosPerPoint         = 1000.0 // $1000.0 = 1 PesosPerPoint
	PesosPerRedeemedPoint = 100.0  // 1 points = $100.0 pesos
)

type Customer struct {
	ID        string    `json:"id" `
	Name      string    `json:"name"`
	Points    int       `json:"points"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Purchase struct {
	ID           string    `json:"id"`
	CustomerID   string    `json:"customer_id"`
	Amount       float64   `json:"amount"`
	EarnedPoints int       `json:"earned_points"`
	CreatedAt    time.Time `json:"created_at"`
}

type Reedemption struct {
	ID             string    `json:"id"`
	CustomerID     string    `json:"customer_id"`
	RedeemedPoints int       `json:"redeemed_points"`
	EquivalentCash float64   `json:"equivalent_cash"`
	CreatedAt      time.Time `json:"created_at"`
}
