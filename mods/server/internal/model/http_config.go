package model

type HttpConfig struct {
	Port string `koanf:"port"`
}

func (c *HttpConfig) GetPort() string {
	return ":" + c.Port
}
