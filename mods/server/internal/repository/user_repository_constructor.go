package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nexus/internal/entity"
)

func NewUserRepository(db *gorm.DB, logger *zap.Logger) *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository[entity.User](db, logger),
	}
}
