package badger_bucket

import (
	"bytes"
	"encoding/gob"
)

func gobMarshal[T any](v T) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// gob 反序列化
func gobUnmarshal[T any](data []byte, v *T) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(v)
}
