package repository

import (
	"nexus/internal/entity"
)

// UserRepository 定义了一个 User 的 Repository 结构
type UserRepository struct {
	*BaseRepository[entity.User]
}

// 特定的 UserRepository 方法
func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
