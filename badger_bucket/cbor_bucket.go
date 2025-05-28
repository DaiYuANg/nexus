package badger_bucket

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/fxamacker/cbor/v2"
)

func NewCborSerializeBucket[T any](db *badger.DB, prefix string) *SerializeBucket[T] {
	return NewSerializeBucket(db, prefix,
		WithMarshal(
			func(v T) ([]byte, error) {
				return cbor.Marshal(v)
			},
		),
		WithUnmarshal(
			func(data []byte, v *T) error {
				return cbor.Unmarshal(data, v)
			},
		),
	)
}
