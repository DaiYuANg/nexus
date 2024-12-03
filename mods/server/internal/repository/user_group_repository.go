package repository

import "nexus/internal/entity"

type UserGroupRepository struct {
	*BaseRepository[entity.UserGroup]
}
