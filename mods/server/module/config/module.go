package config

import (
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"nexus/internal/conf"
	"strings"
)

const EnvPrefix = "NEXUS_"

var Module = fx.Module("config_module", fx.Provide(newConfigParser, parseConfig))

func newConfigParser() *koanf.Koanf {
	kc := koanf.Conf{
		Delim:       ".",
		StrictMerge: true,
	}
	return koanf.NewWithConf(kc)
}

type ParseConfigResult struct {
	fx.Out
	Config         *conf.Config
	DatabaseConfig *conf.DatabaseConfig
	HttpConfig     *conf.HttpConfig
	FileConfig     *conf.FileConfig
	LoggerConfig   *conf.LoggingConfig
}
type ParseParams struct {
	fx.In
	K *koanf.Koanf
}

func parseConfig(params ParseParams) (ParseConfigResult, error) {
	k := params.K
	c, err := defaultConfig()
	if err != nil {
		return ParseConfigResult{}, nil
	}
	lo.Must0(k.Load(structs.Provider(c, "default"), nil))
	lo.Must0(k.Load(env.Provider(EnvPrefix, ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, EnvPrefix)), "_", ".", -1)
	}), nil))

	var out conf.Config
	lo.Must0(k.Unmarshal("", &out), "error unmarshalling config")
	config := &out
	return ParseConfigResult{
		Config:         config,
		DatabaseConfig: &config.Database,
		HttpConfig:     &config.Http,
		LoggerConfig:   &config.Logging,
		FileConfig:     &config.File,
	}, nil
}
