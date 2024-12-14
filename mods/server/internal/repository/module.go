package repository

import (
	"go.uber.org/fx"
)

var Module = fx.Module("repository",
	fx.Provide(
		NewUserRepository,
		NewFileResourceRepository,
		NewFolderRepository,
		NewFileRepository,
		NewUserGroupRepository,
	),
)
