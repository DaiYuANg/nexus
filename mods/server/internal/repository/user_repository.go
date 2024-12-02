package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nexus/internal/entity"
)

// UserRepository 定义了一个 User 的 Repository 结构
type UserRepository struct {
	*BaseRepository[entity.User]
}

// NewUserRepository 创建一个新的 UserRepository 实例
func NewUserRepository(db *gorm.DB, logger *zap.Logger) *UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository[entity.User](db, logger),
	}
}

// 特定的 UserRepository 方法
func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
