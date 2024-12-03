package conf

type MinioConfig struct {
	Endpoint  string `json:"endpoint" koanf:"endpoint"`
	AccessKey string `json:"accessKey" koanf:"accesskey"`
	SecretKey string `json:"secretKey" koanf:"secretkey"`
	UseSSL    bool   `json:"use_ssl" koanf:"use.ssl"`
}
