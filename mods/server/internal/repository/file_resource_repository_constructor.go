package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nexus/internal/entity"
)

func NewFileResourceRepository(db *gorm.DB, logger *zap.Logger) *FileResourceRepository {
	return &FileResourceRepository{
		BaseRepository: NewBaseRepository[entity.FileResource](db, logger),
	}
}
