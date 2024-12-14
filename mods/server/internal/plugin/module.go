package plugin

import (
	"github.com/hashicorp/go-plugin"
	"go.uber.org/fx"
)

var Module = fx.Module("plugin", fx.Provide(createPluginClient))

func createPluginClient() *plugin.Client {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			MagicCookieKey:   "example",
			MagicCookieValue: "12345",
		},
	})
	return client
}
