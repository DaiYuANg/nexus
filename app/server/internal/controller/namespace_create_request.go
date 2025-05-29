package controller

import "github.com/DaiYuANg/maxio/server/internal/bucket"

type NamespaceCreateRequest struct {
	Name           string                    `json:"name" validate:"required"`
	Description    string                    `json:"description"`
	AllowDuplicate bool                      `json:"allow_duplicate"`
	OwnerID        string                    `json:"owner_id"`
	QuotaBytes     int64                     `json:"quota_bytes"`
	AccessControl  string                    `json:"access_control"`
	Tags           []string                  `json:"tags"`
	StorageBackend bucket.StorageBackendType `json:"storage_backend"`
	StorageConfig  map[string]string         `json:"storage_config"`
	StorageMode    bucket.StorageMode        `json:"storage_mode"`
}
