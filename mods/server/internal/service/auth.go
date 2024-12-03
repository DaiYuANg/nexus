package service

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"nexus/internal/error_code"
	"nexus/internal/model"
	"nexus/internal/repository"
)

type Auth struct {
	*JWT
	*zap.Logger
	*repository.UserRepository
}

func (s *Auth) Login(loginUser model.LoginUser) (*model.UserVerified, error) {
	user, err := s.FindByEmail(loginUser.Email)
	if err != nil {
		s.Error("FindByEmail", zap.Error(err))
		return nil, errors.New(error_code.UserNotFound)
	}

	s.Info("User", zap.Any("user", user))

	sign, err := s.Sign(user.Email)
	if err != nil {
		s.Error("Sign", zap.Error(err))
		return nil, err
	}

	verified := model.UserVerified{
		Token:  sign,
		Email:  user.Email,
		Avatar: user.Avatar,
	}
	s.Info("sign", zap.Any("sign", verified))
	return &verified, nil
}
