package config

import (
	"github.com/adrg/xdg"
	"path"
)

func defaultConfig() (Config, error) {
	defaultDataDir, err := xdg.DataFile("nexus_data")
	if err != nil {
		return Config{}, err
	}
	return Config{
		Http: HttpConfig{Port: "3000"},
		Database: DatabaseConfig{
			Type: "sqlite",
			Path: path.Join(defaultDataDir, "nexus.db"),
		},
		File: FileConfig{
			Data: defaultDataDir,
		},
	}, nil
}
