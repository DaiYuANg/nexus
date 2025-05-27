package metadata

import (
	"github.com/influxdata/influxdb/pkg/snowflake"
	"go.uber.org/fx"
)

var Module = fx.Module("metadata", fx.Provide(newDefaultStore, snowflakeId))

func newDefaultStore() (*BboltStore, error) {
	return NewBboltStore("metadata.db")
}

func snowflakeId() *snowflake.Generator {
	return snowflake.New(0)
}
