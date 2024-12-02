package io

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

// CalculateMD5File 计算文件的 MD5 值
func CalculateMD5File(file io.Reader) (string, error) {
	hash := md5.New()

	// 计算 MD5
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("failed to calculate MD5: %w", err)
	}

	// 返回 MD5 值的十六进制表示
	return hex.EncodeToString(hash.Sum(nil)), nil
}
