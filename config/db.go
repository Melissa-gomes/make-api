package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func NewConnDB(env Env) *gorm.DB {
	config := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s",
		env.DB_USERNAME,
		env.DB_PASSWORD,
		env.DB_HOST,
		env.DB_PORT,
		env.DB_DATABASE,
		env.DB_SSLMODE,
	)

	db, err := gorm.Open(postgres.Open(config))
	if err != nil {
		panic("fail to connect in database")
	}

	return db
}
