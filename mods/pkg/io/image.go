package io

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
)

func ImageToBase64(img image.Image) (string, error) {
	// 创建一个 buffer 用于存放图片数据
	var buf bytes.Buffer

	// 将图片编码为 PNG 格式
	err := png.Encode(&buf, img)
	if err != nil {
		return "", err
	}

	// 将 buffer 中的字节数据转换为 Base64 字符串
	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	return base64Str, nil
}
