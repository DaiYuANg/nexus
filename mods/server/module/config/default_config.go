package config

import (
	"github.com/adrg/xdg"
	"nexus/internal/conf"
	"path"
)

func defaultConfig() (conf.Config, error) {
	defaultDataDir, err := xdg.DataFile("nexus_data")
	if err != nil {
		return conf.Config{}, err
	}
	return conf.Config{
		Http: conf.HttpConfig{Port: "3000"},
		Database: conf.DatabaseConfig{
			Type: "sqlite",
			Path: path.Join(defaultDataDir, "nexus.db"),
		},
		File: conf.FileConfig{
			Data: defaultDataDir,
		},
	}, nil
}
