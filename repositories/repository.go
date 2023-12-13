package repositories

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func NewRepository(database *gorm.DB) Repository {
	return Repository{DB: database}
}
