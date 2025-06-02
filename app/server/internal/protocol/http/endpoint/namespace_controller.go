package endpoint

import (
	"github.com/DaiYuANg/maxio/server/internal/bucket"
	"github.com/gofiber/fiber/v2"
)

type NamespaceController struct {
	service *bucket.Service
}

func (n NamespaceController) RegisterRoutes(app *fiber.App) {
	group := app.Group("bucket")
	group.Post("create", n.createNamespace)
	group.Get("list", n.ListNamespace)
}

func (n NamespaceController) createNamespace(ctx *fiber.Ctx) error {
	var req NamespaceCreateRequest
	if err := ctx.BodyParser(&req); err != nil {
		return Fail(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	ns := bucket.New(req.Name,
		bucket.WithQuota(req.QuotaBytes),
		bucket.WithStorage(req.StorageBackend, req.StorageConfig),
		func(n *bucket.Bucket) {
			n.Description = req.Description
			n.OwnerID = req.OwnerID
			n.AllowDuplicate = req.AllowDuplicate
			n.AccessControl = req.AccessControl
			n.Tags = req.Tags
			n.StorageMode = req.StorageMode
		},
	)

	if err := n.service.CreateNamespace(ns); err != nil {
		return Fail(ctx, fiber.StatusConflict, "bucket already exists")
	}

	return Success(ctx, ns)
}

func (n NamespaceController) ListNamespace(ctx *fiber.Ctx) error {
	namespaces, err := n.service.ListNamespaces()
	if err != nil {
		return Fail(ctx, fiber.StatusInternalServerError, err.Error())
	}
	return Success(ctx, namespaces)
}
