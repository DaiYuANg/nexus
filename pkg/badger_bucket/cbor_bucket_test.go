package badger_bucket

import (
	"github.com/dgraph-io/badger/v4"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

// 测试用的结构体
type TestStruct struct {
	ID   int
	Name string
}

func TestNewCborSerializeBucket(t *testing.T) {
	// 临时目录做测试用 DB 路径
	tmp := os.TempDir()
	dir := path.Join(tmp, "test_badger_db")
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Error(err)
		}
	}(dir)

	opts := badger.DefaultOptions(dir).WithLogger(nil) // 关闭日志方便测试
	db, err := badger.Open(opts)
	require.NoError(t, err)
	defer func(db *badger.DB) {
		err := db.Close()
		if err != nil {
			t.Error(err)
		}
	}(db)

	bucket := NewCborSerializeBucket[TestStruct](db, "testprefix")

	// 测试写入
	err = bucket.Put("key1", TestStruct{ID: 123, Name: "Alice"})
	require.NoError(t, err)

	// 测试读取
	val, err := bucket.Get("key1")
	require.NoError(t, err)
	require.Equal(t, 123, val.ID)
	require.Equal(t, "Alice", val.Name)

	// 测试读取不存在 key
	_, err = bucket.Get("nonexistent")
	require.Error(t, err)

	// 测试 Scan
	results := make(map[string]TestStruct)
	err = bucket.Scan(func(key string, val TestStruct) error {
		results[key] = val
		return nil
	})
	require.NoError(t, err)
	require.Contains(t, results, "key1")
	require.Equal(t, "Alice", results["key1"].Name)
}
