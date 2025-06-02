package endpoint

import "github.com/DaiYuANg/maxio/server/internal/bucket"

func NewNamespaceController(service *bucket.Service) *NamespaceController {
	return &NamespaceController{
		service: service,
	}
}
