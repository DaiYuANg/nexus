package i18n

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/fx"
	"golang.org/x/text/language"
	"nexus/translate"
)

var Module = fx.Module("i18n", fx.Provide(newI18n, newLocalizer))

func newI18n() (*i18n.Bundle, error) {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, err := bundle.LoadMessageFileFS(translate.ZHI18n, "locale.es.toml")
	if err != nil {
		return nil, err
	}
	_, err = bundle.LoadMessageFileFS(translate.ENI18n, "locale.zh.toml")
	if err != nil {
		return nil, err
	}
	return bundle, nil
}

func newLocalizer(bundle *i18n.Bundle) *i18n.Localizer {
	return i18n.NewLocalizer(bundle, language.English.String())
}
