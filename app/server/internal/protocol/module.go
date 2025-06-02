package protocol

import (
	"github.com/DaiYuANg/maxio/server/internal/protocol/http"
	"github.com/DaiYuANg/maxio/server/internal/protocol/tcp"
	"go.uber.org/fx"
)

var Module = fx.Module("protocol", tcp.Module, http.Module)
