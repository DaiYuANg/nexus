package jwt

import (
	"crypto/rand"
	"go.uber.org/fx"
)

var Module = fx.Module("jwt",
	fx.Provide(
		fx.Annotate(generateSigningKey,
			fx.ResultTags(`name:"jwtKey"`),
		),
	),
)

func generateSigningKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}
