package repository

import "nexus/internal/entity"

// UserRepository 定义了一个 User 的 Repository 结构
type FileResourceRepository struct {
	*BaseRepository[entity.FileResource]
}
