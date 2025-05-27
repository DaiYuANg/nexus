package internal_store

type Store interface {
	Put(bucket string, key []byte, value []byte) error
	Get(bucket string, key []byte) ([]byte, error)
	Delete(bucket string, key []byte) error
	ListKeys(bucket string) ([][]byte, error) // 可选，列出所有 key
}
