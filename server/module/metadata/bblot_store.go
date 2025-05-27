package metadata

import "go.etcd.io/bbolt"

import (
	"errors"
)

type BboltStore struct {
	db *bbolt.DB
}

// NewBboltStore 打开或创建 bbolt 文件数据库
func NewBboltStore(path string) (*BboltStore, error) {
	db, err := bbolt.Open(path, 0666, nil)
	if err != nil {
		return nil, err
	}
	return &BboltStore{db: db}, nil
}

func (s *BboltStore) Put(bucket string, key []byte, value []byte) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		return b.Put(key, value)
	})
}

func (s *BboltStore) Get(bucket string, key []byte) ([]byte, error) {
	var val []byte
	err := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return errors.New("bucket not found")
		}
		v := b.Get(key)
		if v == nil {
			return errors.New("key not found")
		}
		// 复制数据，防止事务结束后指针失效
		val = make([]byte, len(v))
		copy(val, v)
		return nil
	})
	return val, err
}

func (s *BboltStore) Delete(bucket string, key []byte) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return errors.New("bucket not found")
		}
		return b.Delete(key)
	})
}

func (s *BboltStore) ListKeys(bucket string) ([][]byte, error) {
	var keys [][]byte
	err := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return errors.New("bucket not found")
		}
		return b.ForEach(func(k, v []byte) error {
			kCopy := make([]byte, len(k))
			copy(kCopy, k)
			keys = append(keys, kCopy)
			return nil
		})
	})
	return keys, err
}
