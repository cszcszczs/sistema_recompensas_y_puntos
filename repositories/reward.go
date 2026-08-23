package repositories

import (
	"sync"
	"time"

	"git.com/api-rest/models"
)

// interfaz para definir los metodos que debe cumplir la capa de datos
type RewardRepository interface {
	GetCustomerByID(id string) (*models.Customer, error)
	SaveCustomer(customer *models.Customer) error
	SavePurchase(purchase *models.Purchase) error
	SaveRedemption(redemption *models.Redemption) error
}

// implementacion de la memoria temporal
type MemoryRepository struct {
	mu          sync.RWMutex
	customers   map[string]*models.Customer
	purchases   []*models.Purchase
	redemptions []*models.Redemption
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		customers:   make(map[string]*models.Customer),
		purchases:   make([]*models.Purchase, 0),
		redemptions: make([]*models.Redemption, 0),
	}
}

func (r *MemoryRepository) GetCustomerByID(id string) (*models.Customer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	customer, exists := r.customers[id]
	if !exists {
		return nil, models.ErrCustomerNotFound
	}

	customerCopy := *customer
	return &customerCopy, nil
}

func (r *MemoryRepository) SaveCustomer(customer *models.Customer) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	customer.UpdatedAt = time.Now()
	r.customers[customer.ID] = customer
	return nil
}

func (r *MemoryRepository) SavePurchase(purchase *models.Purchase) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	purchase.CreatedAt = time.Now()
	r.purchases = append(r.purchases, purchase)
	return nil
}

func (r *MemoryRepository) SaveRedemption(redemption *models.Redemption) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	redemption.CreatedAt = time.Now()
	r.redemptions = append(r.redemptions, redemption)
	return nil
}
