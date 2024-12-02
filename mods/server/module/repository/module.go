package repository

import (
	"go.uber.org/fx"
	"nexus/internal/repository"
)

var Module = fx.Module("repository", fx.Provide(repository.NewUserRepository))
