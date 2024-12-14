package ws

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type LiveNotificationParam struct {
	fx.In
	*zap.Logger
}

func NewLiveNotification(param LiveNotificationParam) *LiveNotification {
	return &LiveNotification{
		Logger: param.Logger,
	}
}
