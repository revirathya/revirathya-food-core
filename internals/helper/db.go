package helper

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/revirathya/revirathya-food-core/internals/config"
)

func NewDatabase(C *config.Config) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		C.DBUsername,
		C.DBPassword,
		C.DBHostname,
		C.DBPort,
		C.DBName,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
