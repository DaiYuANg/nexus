package service

import (
	"github.com/jinzhu/copier"
	goeventbus "github.com/stanipetrosyan/go-eventbus"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nexus/internal/constant"
	"nexus/internal/entity"
	"nexus/internal/model"
)

type User struct {
	db       *gorm.DB
	logger   *zap.Logger
	eventbus goeventbus.EventBus
}

func (s *User) Register(registerUser model.RegisterUser) {
	user := entity.User{}
	err := copier.Copy(&user, &registerUser)
	s.logger.Info("Copied", zap.Any("registerUser", user))
	if err != nil {
		return
	}

	s.db.Create(&user)
	options := goeventbus.NewMessageOptions().SetHeaders(goeventbus.Headers{})
	message := goeventbus.CreateMessage().SetOptions(options).SetBody(user)
	s.eventbus.Channel(constant.UserRegistered).Publisher().Publish(message)
}

type UserServiceParam struct {
	fx.In
	EventBus goeventbus.EventBus
	DB       *gorm.DB
	Logger   *zap.Logger
}

func NewUserService(param UserServiceParam) *User {
	return &User{
		db:       param.DB,
		logger:   param.Logger,
		eventbus: param.EventBus,
	}
}
