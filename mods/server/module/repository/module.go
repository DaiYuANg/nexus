package repository

import (
	"go.uber.org/fx"
	"nexus/internal/repository"
)

var Module = fx.Module("repository",
	fx.Provide(
		repository.NewUserRepository,
		repository.NewFileResourceRepository,
		repository.NewFolderRepository,
		repository.NewFileRepository,
		repository.NewUserGroupRepository,
	),
)
