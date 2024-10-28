package orders

import (
	"gorm.io/gorm"

	"github.com/revirathya/revirathya-food-core/internals/domain"
)

type OrderDatabaseRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderDatabaseRepository(db *gorm.DB) domain.OrderDatabaseRepository {
	return &OrderDatabaseRepositoryImpl{
		db: db,
	}
}

func (r *OrderDatabaseRepositoryImpl) GetOrderByID(id int) (*domain.Order, error) {
	return &domain.Order{
		ID:       id,
		Date:     "2024-10-28",
		Category: "Breakfast",
		Location: "Wisma 46 - KFC",
		Amount:   20000,
	}, nil
}

func (r *OrderDatabaseRepositoryImpl) GetOrders(date string, category string, location string) ([]domain.Order, error) {
	order := domain.Order{
		ID:       1,
		Date:     date,
		Category: category,
		Location: location,
		Amount:   20000,
	}
	return []domain.Order{order}, nil
}

func (r *OrderDatabaseRepositoryImpl) CreateOrder(date string, category string, location string, amount int) (*domain.Order, error) {
	return &domain.Order{
		ID:       1,
		Date:     date,
		Category: category,
		Location: location,
		Amount:   amount,
	}, nil
}
