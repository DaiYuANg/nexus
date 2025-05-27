package namespace

import (
	"fmt"
	"go.etcd.io/bbolt"
	"go.uber.org/fx"
)

var Module = fx.Module("namespace",
	fx.Provide(
		newNamespaceStore,
	),
)

type Store struct {
	db *bbolt.DB
}

func (s Store) CreateNamespace(name string) error {
	err := s.db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func newNamespaceStore() (*Store, error) {
	db, err := bbolt.Open("namespace.db", 0600, nil)
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}
