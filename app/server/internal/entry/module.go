package entry

import (
	"github.com/DaiYuANg/maxio/server/internal/entry/http"
	"github.com/DaiYuANg/maxio/server/internal/entry/tcp"
	"go.uber.org/fx"
)

var Module = fx.Module("entry", tcp.Module, http.Module)
