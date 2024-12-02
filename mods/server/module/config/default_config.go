package config

import (
	"github.com/adrg/xdg"
	"nexus/internal/model"
	"path"
)

var defaultConfig = model.Config{
	Http: model.HttpConfig{Port: "3000"},
	Database: model.DatabaseConfig{
		Type: "sqlite",
		Path: path.Join(xdg.CacheHome, "nexus.db"),
	},
	Logging: model.LoggingConfig{
		Path: path.Join(xdg.CacheHome, "nexus.log"),
	},
}
