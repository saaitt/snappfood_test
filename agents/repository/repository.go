package repository

import (
	"github.com/saaitt/snappfood_test/agents/domain"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.Repository {
	return &repository{
		db: db,
	}
}
