package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nexus/internal/entity"
)

func NewUserGroupRepository(db *gorm.DB, logger *zap.Logger) *UserGroupRepository {
	return &UserGroupRepository{
		BaseRepository: NewBaseRepository[entity.UserGroup](db, logger),
	}
}
