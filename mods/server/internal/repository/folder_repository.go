package repository

import "nexus/internal/entity"

type FolderRepository struct {
	*BaseRepository[entity.Folder]
}
