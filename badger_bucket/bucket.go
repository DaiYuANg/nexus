package badger_bucket

import (
	"github.com/dgraph-io/badger/v4"
	"strings"
)

// 纯 byte slice 的底层 Bucket
type Bucket struct {
	db     *badger.DB
	prefix string
}

func NewBucket(db *badger.DB, prefix string) *Bucket {
	if !strings.HasSuffix(prefix, ":") {
		prefix += ":"
	}
	return &Bucket{db: db, prefix: prefix}
}

func (b *Bucket) Put(key string, value []byte) error {
	fullKey := []byte(b.prefix + key)
	return b.db.Update(func(txn *badger.Txn) error {
		return txn.Set(fullKey, value)
	})
}

func (b *Bucket) Get(key string) ([]byte, error) {
	fullKey := []byte(b.prefix + key)
	var valCopy []byte
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(fullKey)
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			valCopy = append([]byte{}, val...)
			return nil
		})
	})
	return valCopy, err
}

func (b *Bucket) Scan(fn func(key string, val []byte) error) error {
	prefix := []byte(b.prefix)
	opts := badger.DefaultIteratorOptions
	opts.Prefix = prefix

	return b.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(val []byte) error {
				keyStr := string(k[len(prefix):])
				return fn(keyStr, val)
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
}
