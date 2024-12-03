package service

import (
	"github.com/gomig/avatar"
	"github.com/jinzhu/copier"
	goeventbus "github.com/stanipetrosyan/go-eventbus"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"nexus/internal/constant"
	"nexus/internal/entity"
	"nexus/internal/model"
	"nexus/internal/repository"
)

type User struct {
	db *gorm.DB
	*repository.UserRepository
	*zap.Logger
	goeventbus.EventBus
}

func (s *User) Register(registerUser model.RegisterUser) error {
	user := entity.User{}
	err := copier.Copy(&user, &registerUser)
	s.Info("Copied", zap.Any("registerUser", user))
	if err != nil {
		return err
	}
	john := avatar.NewTextAvatar(registerUser.Email)
	user.Avatar = john.Base64()
	err = s.Create(&user)
	if err != nil {
		return err
	}

	options := goeventbus.NewMessageOptions().SetHeaders(goeventbus.Headers{})
	message := goeventbus.CreateMessage().SetOptions(options).SetBody(user)
	s.Channel(constant.UserRegistered).Publisher().Publish(message)
	return nil
}
