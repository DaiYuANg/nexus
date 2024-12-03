package repository

import "nexus/internal/entity"

type FileRepository struct {
	*BaseRepository[entity.File]
}
