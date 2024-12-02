package route

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
	"nexus/internal/model"
	"nexus/internal/service"
)

type User struct {
	*service.User
	*service.Auth
	*zap.Logger
}

func (r User) register(context *fiber.Ctx) error {
	registerUser := model.RegisterUser{}
	err := context.BodyParser(&registerUser)
	r.Debug("Param", zap.Any("Body", registerUser))
	if err != nil {
		return err
	}
	r.Register(registerUser)
	return context.Status(200).JSON(model.JustOk())
}

func (r User) login(context *fiber.Ctx) error {
	loginUser := model.LoginUser{}
	err := context.BodyParser(&loginUser)
	r.Debug("Param", zap.Any("Body", loginUser))
	if err != nil {
		return err
	}
	jwt, err := r.Login(loginUser)
	if err != nil {
		return err
	}
	return context.Status(200).JSON(model.Ok(jwt))
}

func (r User) Routes() []Info {
	return []Info{
		{
			Path:    "/user/register",
			Method:  http.MethodPost,
			Handler: r.register,
		},
		{
			Path:    "/user/login",
			Method:  http.MethodPost,
			Handler: r.login,
		},
	}
}

func NewUserRoute(userService *service.User, logger *zap.Logger, auth *service.Auth) *User {
	return &User{
		User:   userService,
		Logger: logger,
		Auth:   auth,
	}
}
