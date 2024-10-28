package domain

// Entity
type Order struct {
	ID       int    `json:"id"`
	Date     string `json:"date"`
	Category string `json:"category"`
	Location string `json:"location"`
	Amount   int    `json:"amount"`
}

// Repository
type OrderDatabaseRepository interface {
	GetOrderByID(id int) (*Order, error)
	GetOrders(date string, category string, location string) ([]Order, error)
	CreateOrder(date string, category string, location string, amount int) (*Order, error)
}

// Usercase
type OrderUsecase interface {
	GetOrderByID(id int) (*Order, error)
	GetOrders(date string, category string, location string) ([]Order, error)
	CreateOrder(date string, category string, location string, amount int) (*Order, error)
}
