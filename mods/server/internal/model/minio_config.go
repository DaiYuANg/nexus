package model

type MinioConfig struct {
	Endpoint  string `json:"endpoint" koanf:"endpoint"`
	AccessKey string `json:"access_key" koanf:"access_key"`
	SecretKey string `json:"secret_key" koanf:"secret_key"`
	UseSSL    bool   `json:"use_ssl" koanf:"use_ssl"`
}
