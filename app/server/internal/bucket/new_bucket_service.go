package bucket

import (
	"github.com/samber/lo"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
)

func NewService(db *bbolt.DB, logger *zap.SugaredLogger) *Service {
	lo.Must0(
		db.Update(
			func(tx *bbolt.Tx) error {
				_, err := tx.CreateBucketIfNotExists(bucketKey)
				return err
			},
		),
	)
	return &Service{
		db,
		logger,
	}
}
