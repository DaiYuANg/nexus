package namespace

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"go.uber.org/zap"
	"time"
)

const (
	namespace = "namespace"
)

var (
	namespaceBucket = []byte(namespace)
)

type Service struct {
	*bbolt.DB
	*zap.SugaredLogger
}

// CreateNamespace Create 创建一个新的命名空间（namespace），名称必须唯一
func (s Service) CreateNamespace(ns *Namespace) error {
	s.Debugw("creating namespace", zap.String("name", ns.Name))

	return s.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(namespaceBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", namespace)
		}

		// 可选去重校验
		if bucket.Get([]byte(ns.Name)) != nil && !ns.AllowDuplicate {
			return fmt.Errorf("namespace %q already exists", ns.Name)
		}

		ns.UpdatedAt = time.Now()

		data, err := json.Marshal(ns)
		if err != nil {
			return fmt.Errorf("marshal namespace: %w", err)
		}
		return bucket.Put([]byte(ns.Name), data)
	})
}

func (s Service) ListNamespaces() ([]*Namespace, error) {
	var result []*Namespace

	err := s.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket(namespaceBucket)
		if bucket == nil {
			return fmt.Errorf("bucket %q not found", namespace)
		}

		return bucket.ForEach(func(k, v []byte) error {
			var ns Namespace
			if err := json.Unmarshal(v, &ns); err != nil {
				s.Warnw("failed to unmarshal namespace", zap.String("key", string(k)), zap.Error(err))
				// 可选：忽略损坏数据 or return err
				return nil
			}
			result = append(result, &ns)
			return nil
		})
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s Service) NamespaceExists(name string) bool {
	var exists bool
	err := s.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(namespaceBucket)
		v := b.Get([]byte(name))
		exists = v != nil
		return nil
	})
	if err != nil {
		s.Warn("failed to check if namespace exists", zap.String("name", name), zap.Error(err))
		return false
	}
	return exists
}
