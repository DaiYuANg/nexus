package model

type Config struct {
	Http     HttpConfig     `koanf:"http"`
	Database DatabaseConfig `koanf:"database"`
	Minio    MinioConfig    `koanf:"minio"`
	Logging  LoggingConfig  `koanf:"logging"`
	Email    EmailConfig    `koanf:"email"`
}

type LoggingConfig struct {
	Path string `koanf:"path"`
}
