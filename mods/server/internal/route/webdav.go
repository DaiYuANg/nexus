package route

import (
	"net/http"
)

const webdavRoot = "./webdav_storage"

type Webdav struct {
}

// WebDAV 认证处理
func basicAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 基本认证
		user, password, ok := r.BasicAuth()
		if !ok || user != "admin" || password != "password" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		handler.ServeHTTP(w, r)
	})
}

func (r Webdav) Routes() []Info {
	return []Info{
		{
			Path:   "/runtime/info",
			Method: ALL,
		},
	}
}
