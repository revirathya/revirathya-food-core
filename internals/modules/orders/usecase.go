package orders

import (
	"github.com/revirathya/revirathya-food-core/internals/domain"
)

type OrderUsecaseImpl struct {
	dbRepository domain.OrderDatabaseRepository
}

func NewOrderUsecase(dbRepository domain.OrderDatabaseRepository) domain.OrderUsecase {
	return &OrderUsecaseImpl{
		dbRepository: dbRepository,
	}
}

func (u *OrderUsecaseImpl) GetOrderByID(id int) (*domain.Order, error) {
	return u.dbRepository.GetOrderByID(id)
}

func (u *OrderUsecaseImpl) GetOrders(date string, category string, location string) ([]domain.Order, error) {
	return u.dbRepository.GetOrders(date, category, location)
}

func (u *OrderUsecaseImpl) CreateOrder(date string, category string, location string, amount int) (*domain.Order, error) {
	return u.dbRepository.CreateOrder(date, category, location, amount)
}
