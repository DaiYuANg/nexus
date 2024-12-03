package service

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type JwtServiceParam struct {
	fx.In
	SigningKey []byte `name:"jwtKey"`
	*zap.Logger
}

func NewJWTService(param JwtServiceParam) *JWT {
	return &JWT{
		Logger:     param.Logger,
		signingKey: param.SigningKey,
	}
}
