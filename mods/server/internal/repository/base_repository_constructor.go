package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewBaseRepository[T any](db *gorm.DB, logger *zap.Logger) *BaseRepository[T] {
	return &BaseRepository[T]{DB: db, Logger: logger}
}
