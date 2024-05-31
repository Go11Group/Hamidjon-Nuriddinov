package packages

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("postgres://postgres:hamidjon4424@localhost:5432/mybase?sslmode=disable"))
	if err != nil {
		return nil, err
	}
	return db, nil
}
