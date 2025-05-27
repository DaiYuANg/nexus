package controller

import "github.com/DaiYuANg/storix/server/internal/namespace"

type NamespaceCreateRequest struct {
	Name           string                       `json:"name" validate:"required"`
	Description    string                       `json:"description"`
	AllowDuplicate bool                         `json:"allow_duplicate"`
	OwnerID        string                       `json:"owner_id"`
	QuotaBytes     int64                        `json:"quota_bytes"`
	AccessControl  string                       `json:"access_control"`
	Tags           []string                     `json:"tags"`
	StorageBackend namespace.StorageBackendType `json:"storage_backend"`
	StorageConfig  map[string]string            `json:"storage_config"`
	StorageMode    namespace.StorageMode        `json:"storage_mode"`
}
