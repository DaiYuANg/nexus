package route

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"nexus/pkg/runtime"
)

type Runtime struct {
}

func (r Runtime) systemInfo(context *fiber.Ctx) error {
	return context.Status(200).JSON(runtime.BuildRuntimeInfo())
}

func (r Runtime) Routes() []Info {
	return []Info{
		{
			Path:    "/runtime/info",
			Method:  http.MethodGet,
			Handler: r.systemInfo,
		},
	}
}

func NewRuntimeRoute() *Runtime {
	return &Runtime{}
}
