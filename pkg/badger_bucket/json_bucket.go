package badger_bucket

import (
	"encoding/json"
	"github.com/dgraph-io/badger/v4"
)

func NewJsonSerializeBucket[T any](db *badger.DB, prefix string) *SerializeBucket[T] {
	return NewSerializeBucket(db, prefix,
		WithMarshal(
			func(v T) ([]byte, error) {
				return json.Marshal(v)
			},
		),
		WithUnmarshal(
			func(data []byte, v *T) error {
				return json.Unmarshal(data, v)
			},
		),
	)
}
