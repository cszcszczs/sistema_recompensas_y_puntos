package services

import (
	"math"

	"git.com/api-rest/models"
	"git.com/api-rest/repositories"
)

type RewardService interface {
	RegisterPurchase(req models.PurchaseRequest) (*models.PurchaseResponse, error)
	GetCustomerPoints(customerID string) (*models.PointsBalanceResponse, error)
	RedeemPoints(customerID string, req models.RedeemRequest) (*models.RedeemResponse, error)
}

type rewardService struct {
	repo repositories.RewardRepository
}

func NewRewardService(repo repositories.RewardRepository) RewardService {
	return &rewardService{repo: repo}
}

func (s *rewardService) RegisterPurchase(req models.PurchaseRequest) (*models.PurchaseResponse, error) {
	if req.Amount <= 0 {
		return nil, models.ErrInvalidAmount
	}

	earnedPoints := int(math.Floor(req.Amount / models.PesosPerPoint))

	customer, err := s.repo.GetCustomerByID(req.CustomerID)
	if err != nil {
		if err == models.ErrCustomerNotFound {
			customer = &models.Customer{
				ID:     req.CustomerID,
				Name:   "Customer" + req.CustomerID,
				Points: 0,
			}
		} else {
			return nil, err
		}
	}

	customer.Points += earnedPoints

	if err := s.repo.SaveCustomer(customer); err != nil {
		return nil, err
	}

	purchase := &models.Purchase{
		ID:           "pur_" + req.CustomerID,
		CustomerID:   req.CustomerID,
		Amount:       req.Amount,
		EarnedPoints: earnedPoints,
	}

	if err := s.repo.SavePurchase(purchase); err != nil {
		return nil, err
	}

	return &models.PurchaseResponse{
		PurchaseID:   purchase.ID,
		CustomerID:   customer.ID,
		Amount:       req.Amount,
		EarnedPoints: earnedPoints,
		TotalPoints:  customer.Points,
	}, nil
}

func (s *rewardService) GetCustomerPoints(customerID string) (*models.PointsBalanceResponse, error) {
	customer, err := s.repo.GetCustomerByID(customerID)
	if err != nil {
		return nil, err
	}

	return &models.PointsBalanceResponse{
		CustomerID: customerID,
		Points:     customer.Points,
	}, nil
}

func (s *rewardService) RedeemPoints(customerID string, req models.RedeemRequest) (*models.RedeemResponse, error) {
	if req.Points <= 0 {
		return nil, models.ErrInvalidPoints
	}

	customer, err := s.repo.GetCustomerByID(customerID)
	if err != nil {
		return nil, err
	}

	if customer.Points < req.Points {
		return nil, models.ErrInsufficientPoints
	}

	equivalentCash := float64(req.Points) * models.PesosPerRedeemedPoint

	customer.Points -= req.Points

	if err := s.repo.SaveCustomer(customer); err != nil {
		return nil, err
	}

	redemption := &models.Redemption{
		ID:             "red_" + customerID,
		CustomerID:     customerID,
		RedeemedPoints: req.Points,
		EquivalentCash: equivalentCash,
	}

	if err := s.repo.SaveRedemption(redemption); err != nil {
		return nil, err
	}

	return &models.RedeemResponse{
		RedemptionID:    redemption.ID,
		CustomerID:      customerID,
		RedeemedPoints:  req.Points,
		EquivalentCash:  equivalentCash,
		RemainingPoints: customer.Points,
	}, nil
}
