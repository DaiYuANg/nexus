package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// BaseRepository 定义了一个基础的 Repository 结构
type BaseRepository[T any] struct {
	*gorm.DB
	*zap.Logger
}

// NewBaseRepository 创建一个新的 BaseRepository 实例
func NewBaseRepository[T any](db *gorm.DB, logger *zap.Logger) *BaseRepository[T] {
	return &BaseRepository[T]{DB: db, Logger: logger}
}

// Create 用于创建一个新的实体
func (r *BaseRepository[T]) Create(entity *T) error {
	r.Info("Create", zap.Any("entity", entity))
	if err := r.DB.Create(entity).Error; err != nil {
		return err
	}
	return nil
}

// FindByID 根据 ID 查找实体
func (r *BaseRepository[T]) FindByID(id uint) (*T, error) {
	var entity T
	if err := r.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// FindAll 获取所有实体
func (r *BaseRepository[T]) FindAll() ([]T, error) {
	var entities []T
	if err := r.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

// Update 更新实体
func (r *BaseRepository[T]) Update(entity *T) error {
	if err := r.Save(entity).Error; err != nil {
		return err
	}
	return nil
}

// Delete 根据 ID 删除实体
func (r *BaseRepository[T]) Delete(entity *T) error {
	if err := r.DB.Delete(entity).Error; err != nil {
		return err
	}
	return nil
}
