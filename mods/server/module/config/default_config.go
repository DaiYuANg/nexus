package config

import (
	"github.com/adrg/xdg"
	"github.com/samber/lo"
	"nexus/internal/conf"
	"os"
	"path"
)

func defaultConfig() conf.Config {
	pwd := lo.Must(os.Getwd())
	return conf.Config{
		Http: conf.HttpConfig{Port: "3000"},
		Database: conf.DatabaseConfig{
			Type: "sqlite",
			Path: path.Join(xdg.CacheHome, "nexus.db"),
		},
		Logging: conf.LoggingConfig{
			Path: path.Join(xdg.CacheHome, "nexus.log"),
		},
		File: conf.FileConfig{
			Path: path.Join(pwd, "nexus"),
		},
	}
}
