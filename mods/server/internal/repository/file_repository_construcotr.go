package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nexus/internal/entity"
)

func NewFileRepository(db *gorm.DB, logger *zap.Logger) *FileRepository {
	return &FileRepository{
		BaseRepository: NewBaseRepository[entity.File](db, logger),
	}
}
