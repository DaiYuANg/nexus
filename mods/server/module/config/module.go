package config

import (
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"nexus/internal/model"
	"strings"
)

const EnvPrefix = "NEXUS_"

var Module = fx.Module("config_module", fx.Provide(newConfigParser, parseConfig))

func newConfigParser() *koanf.Koanf {
	conf := koanf.Conf{
		Delim:       ".",
		StrictMerge: true,
	}
	return koanf.NewWithConf(conf)
}

type ParseConfigResult struct {
	fx.Out
	Config         *model.Config
	DatabaseConfig *model.DatabaseConfig
	HttpConfig     *model.HttpConfig
	MinioConfig    *model.MinioConfig
	LoggerConfig   *model.LoggingConfig
}
type ParseParams struct {
	fx.In
	K *koanf.Koanf
}

func parseConfig(params ParseParams) ParseConfigResult {
	k := params.K
	lo.Must0(k.Load(structs.Provider(defaultConfig, "default"), nil))
	lo.Must0(k.Load(env.Provider(EnvPrefix, ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, EnvPrefix)), "_", ".", -1)
	}), nil))
	var out model.Config
	lo.Must0(k.Unmarshal("", &out), "error unmarshalling config")
	config := &out
	return ParseConfigResult{
		Config:         config,
		DatabaseConfig: &config.Database,
		HttpConfig:     &config.Http,
		LoggerConfig:   &config.Logging,
		MinioConfig:    &config.Minio,
	}
}
