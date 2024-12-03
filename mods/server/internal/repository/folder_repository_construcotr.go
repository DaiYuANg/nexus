package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nexus/internal/entity"
)

func NewFolderRepository(db *gorm.DB, logger *zap.Logger) *FolderRepository {
	return &FolderRepository{
		BaseRepository: NewBaseRepository[entity.Folder](db, logger),
	}
}
