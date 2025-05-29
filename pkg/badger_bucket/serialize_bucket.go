package badger_bucket

import (
	"github.com/dgraph-io/badger/v4"
)

type SerializeBucket[T any] struct {
	bucket    *Bucket
	marshal   func(v T) ([]byte, error)
	unmarshal func(data []byte, v *T) error
}

type Option[T any] func(*SerializeBucket[T])

// WithMarshal 配置 marshal 函数
func WithMarshal[T any](marshal func(v T) ([]byte, error)) Option[T] {
	return func(sb *SerializeBucket[T]) {
		sb.marshal = marshal
	}
}

// WithUnmarshal 配置 unmarshal 函数
func WithUnmarshal[T any](unmarshal func(data []byte, v *T) error) Option[T] {
	return func(sb *SerializeBucket[T]) {
		sb.unmarshal = unmarshal
	}
}

func (s *SerializeBucket[T]) Put(key string, val T) error {
	data, err := s.marshal(val)
	if err != nil {
		return err
	}
	return s.bucket.Put(key, data)
}

func (s *SerializeBucket[T]) Get(key string) (T, error) {
	var zero T
	data, err := s.bucket.Get(key)
	if err != nil {
		return zero, err
	}
	var val T
	if err := s.unmarshal(data, &val); err != nil {
		return zero, err
	}
	return val, nil
}

func (s *SerializeBucket[T]) Scan(fn func(key string, val T) error) error {
	return s.bucket.Scan(func(key string, data []byte) error {
		var val T
		if err := s.unmarshal(data, &val); err != nil {
			return err
		}
		return fn(key, val)
	})
}

// NewSerializeBucket 构造函数，接受可变 Option
func NewSerializeBucket[T any](db *badger.DB, prefix string, opts ...Option[T]) *SerializeBucket[T] {
	sb := &SerializeBucket[T]{
		bucket:    NewBucket(db, prefix),
		marshal:   gobMarshal[T],
		unmarshal: gobUnmarshal[T],
	}

	for _, opt := range opts {
		opt(sb)
	}
	return sb
}
