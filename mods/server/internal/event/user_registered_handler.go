package event

import (
	goeventbus "github.com/stanipetrosyan/go-eventbus"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"nexus/internal/constant"
	"nexus/internal/entity"
	"nexus/internal/minio"
)

type UserRegisteredConsumerParam struct {
	fx.In
	goeventbus.EventBus
	*zap.Logger
	*minio.Wrapper
}

func UserRegisteredConsumer(param UserRegisteredConsumerParam) {
	param.Channel(constant.UserRegistered).Subscriber().Listen(func(context goeventbus.Context) {
		user := context.Result().Data.(entity.User)
		param.Info("Register", zap.Any("user", user))
		err := param.CreateBucket(user.Email)
		if err != nil {
			param.Error("CreateBucket", zap.Error(err))
			return
		}
	})
}
