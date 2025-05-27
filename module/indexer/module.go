package indexer

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
	"go.uber.org/fx"
)

var Module = fx.Module("indexer", fx.Provide(newBleve))

func newBleve() *mapping.IndexMappingImpl {
	return bleve.NewIndexMapping()
}
