package config

type Config struct {
	Http     HttpConfig     `koanf:"http"`
	Database DatabaseConfig `koanf:"database"`
	Logging  LoggingConfig  `koanf:"logging"`
	Email    EmailConfig    `koanf:"email"`
	File     FileConfig     `koanf:"file"`
}

type LoggingConfig struct {
	Path  string `koanf:"path"`
	Level string `koanf:"level"`
}
