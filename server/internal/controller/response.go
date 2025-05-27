package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/samber/lo"
	"time"
)

type Response[T any] struct {
	Code      int64  `json:"code"`
	Message   string `json:"message"`
	Data      T      `json:"data,omitempty"`
	Timestamp int64  `json:"timestamp"`
	TraceId   string `json:"trace_id,omitempty"`
	RequestId string `json:"request_id,omitempty"`
}

type ResponseOption func(resp *Response[any])

func WithTraceID(traceId string) ResponseOption {
	return func(resp *Response[any]) {
		resp.TraceId = traceId
	}
}

func WithRequestID(requestId string) ResponseOption {
	return func(resp *Response[any]) {
		resp.RequestId = requestId
	}
}

func Success[T any](ctx *fiber.Ctx, data T, opts ...ResponseOption) error {
	return respond(ctx, 0, "success", data, opts...)
}

func Fail(ctx *fiber.Ctx, code int64, message string, opts ...ResponseOption) error {
	return respond[any](ctx, code, message, nil, opts...)
}

func applyOptions[T any](resp *Response[T], opts []ResponseOption) {
	lo.ForEach(opts, func(opt ResponseOption, _ int) {
		opt(any(resp).(*Response[any]))
	})
}

func respond[T any](ctx *fiber.Ctx, code int64, message string, data T, opts ...ResponseOption) error {
	// 获取 requestid，没则生成
	requestId := ctx.Locals("requestid")
	reqIDStr := ""
	if rid, ok := requestId.(string); ok && rid != "" {
		reqIDStr = rid
	} else {
		reqIDStr = generateID()
	}

	resp := Response[T]{
		Code:      code,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().UnixMilli(),
		RequestId: reqIDStr,
	}

	if resp.TraceId == "" {
		resp.TraceId = generateID()
	}

	applyOptions(&resp, opts)

	return ctx.JSON(resp)
}

func generateID() string {
	return uuid.NewString()
}
