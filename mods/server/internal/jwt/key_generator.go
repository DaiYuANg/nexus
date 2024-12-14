package jwt

import (
	"crypto/rand"
	"github.com/spf13/afero"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const nexusSignKey = "nexus_sign_key"

type GenerateDep struct {
	fx.In
	LocalFs afero.Fs `name:"temp"`
	*zap.Logger
}

func getSigningKey(dep GenerateDep) ([]byte, error) {
	state, err := dep.LocalFs.Stat(nexusSignKey)
	if err != nil {
		key, err := generateSigningKey()
		if err != nil {
			return nil, err
		}

		create, err := dep.LocalFs.Create(nexusSignKey)
		if err != nil {
			return nil, err
		}

		defer create.Close()
		_, err = create.Write(key)
		if err != nil {
			return nil, err
		}

		return key, nil
	}
	dep.Info("Key FIle", zap.String("file", state.Name()))
	open, err := dep.LocalFs.Open(nexusSignKey)
	if err != nil {
		return nil, err
	}
	defer open.Close()
	var key = make([]byte, 1)
	_, err = open.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func generateSigningKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}
