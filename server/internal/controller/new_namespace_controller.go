package controller

import "github.com/DaiYuANg/storix/server/internal/namespace"

func NewNamespaceController(service *namespace.Service) *NamespaceController {
	return &NamespaceController{
		service: service,
	}
}
