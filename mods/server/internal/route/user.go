package route

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"net/http"
	"nexus/internal/model"
	"nexus/internal/service"
)

type User struct {
	userService *service.UserService
	logger      *zap.Logger
}

func (r User) register(context *fiber.Ctx) error {
	registerUser := model.RegisterUser{}
	err := context.BodyParser(&registerUser)
	r.logger.Debug("Param", zap.Any("Body", registerUser))
	if err != nil {
		return err
	}
	r.userService.Register(registerUser)
	return context.Status(200).JSON(model.JustOk())
}

func (r User) Routes() []Info {
	return []Info{
		{
			Path:    "/user/register",
			Method:  http.MethodPost,
			Handler: r.register,
		},
	}
}

func NewUserRoute(userService *service.UserService, logger *zap.Logger) *User {
	return &User{
		userService: userService,
		logger:      logger,
	}
}
