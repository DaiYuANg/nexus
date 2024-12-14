package route

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type FileParam struct {
	fx.In
	*zap.Logger
}

func NewFileRoute(param FileParam) *File {
	return &File{
		Logger: param.Logger,
	}
}
