package store

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (store *Postgres) NewStore() error {
	dsn := "host=localhost user=golang12 password=golang dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	} else {
		store.DB = db
	}
	fmt.Printf("store = %v\n", db)

	return nil
}

type StoreOptions interface {
	NewStore() error
}
