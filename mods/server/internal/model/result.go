package model

import (
	"net/http"
	"strconv"
)

type R[T any] struct {
	Code string `json:"code"`
	Data T      `json:"data,omitempty" swaggerignore:"true"`
}

func Ok[T any](data T) R[T] {
	return R[T]{
		Code: strconv.Itoa(http.StatusOK),
		Data: data,
	}
}

func JustOk() R[any] {
	return R[any]{
		Code: strconv.Itoa(http.StatusOK),
	}
}

func Err[T *any]() R[T] {
	return R[T]{
		Code: strconv.Itoa(http.StatusInternalServerError),
		Data: nil,
	}
}
