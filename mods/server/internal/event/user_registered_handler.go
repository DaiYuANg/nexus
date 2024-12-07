package event

import (
	goeventbus "github.com/stanipetrosyan/go-eventbus"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"io/fs"
	"nexus/internal/constant"
	"nexus/internal/entity"
	"nexus/internal/repository"
	"nexus/vfs/local"
)

type UserRegisteredConsumerParam struct {
	fx.In
	goeventbus.EventBus
	*zap.Logger
	//*fs.Wrapper
	*local.VFS
	*repository.FolderRepository
}

type UserConsumer interface {
	Consumer(user entity.User)
}

func UserRegisteredConsumer(param UserRegisteredConsumerParam) {
	param.Channel(constant.UserRegistered).Subscriber().Listen(func(context goeventbus.Context) {
		user := context.Result().Data.(entity.User)
		param.Info("Register", zap.Any("user", user))
		err := param.Mkdir(user.Email, fs.ModeDir)
		if err != nil {
			param.Error("Create Directory", zap.Error(err))
		}
	})
}
