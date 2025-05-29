package config

type Config struct {
}

type Http struct {
	Enable bool   `toml:"enable"`
	Port   int    `toml:"port"`
	Host   string `toml:"host"`
}
