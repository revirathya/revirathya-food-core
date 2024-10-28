package modules

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/revirathya/revirathya-food-core/internals/config"
	"github.com/revirathya/revirathya-food-core/internals/domain"
	"github.com/revirathya/revirathya-food-core/internals/helper"

	// "github.com/revirathya/revirathya-food-core/internals/delivery/graphql"

	"github.com/revirathya/revirathya-food-core/internals/modules/orders"
)

// Repository
type Repository struct {
	OrderDatabaseRepository domain.OrderDatabaseRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		OrderDatabaseRepository: orders.NewOrderDatabaseRepository(db),
	}
}

// Usecase
type Usecase struct {
	OrderUsecase domain.OrderUsecase
}

func NewUsecase(repository *Repository) *Usecase {
	return &Usecase{
		OrderUsecase: orders.NewOrderUsecase(repository.OrderDatabaseRepository),
	}
}

// Bootstrap
func Bootstrap(C *config.Config, app *fiber.App) error {
	db, err := helper.NewDatabase(C)
	if err != nil {
		return err
	}

	repository := NewRepository(db)
	usecase := NewUsecase(repository)

	// graphql.Setup(app, usecase)
}
