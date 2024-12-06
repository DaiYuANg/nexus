package sftp

import (
	"github.com/studio-b12/gowebdav"
	"go.uber.org/zap"
)

type Config struct {
	Root     string
	User     string
	Password string
	Logger   *zap.Logger
}

func NewWebdav(config Config) (*Webdav, error) {
	client := gowebdav.NewClient(config.Root, config.User, config.Password)

	err := client.Connect()
	if err != nil {
		return nil, err
	}

	return &Webdav{
		Logger: config.Logger,
		Client: client,
	}, nil
}
