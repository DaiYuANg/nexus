package route

import "github.com/gofiber/fiber/v2"

type HTTPMethod string

const (
	Tag               = `group:"route"`
	GET    HTTPMethod = fiber.MethodGet
	POST   HTTPMethod = fiber.MethodPost
	PUT    HTTPMethod = fiber.MethodPut
	DELETE HTTPMethod = fiber.MethodDelete
	ALL    HTTPMethod = "ALL"
)

type Info struct {
	Method    HTTPMethod
	Path      string
	Handler   fiber.Handler
	PermitAll bool
}

type Route interface {
	Routes() []Info
}
