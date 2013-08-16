package fileserver

import (
	"net/http"
)

// Serve the directory `dir` using HTTP on the specified TCP address.
//
// `dir`: filepath to serve.
// `address`: TCP address (`"<host>:<port>"`). The host can be omitted.
// `cache`: If false, the browser is sent headers to prevent it from caching
// content.
func ServeDir(dir, address string, cache bool) error {
	fs := NewFileServer(http.Dir(dir), cache)
	return http.ListenAndServe(address, fs)
}

type FileServer struct {
	handler http.Handler
	cache   bool
}

func NewFileServer(dir http.FileSystem, cache bool) *FileServer {
	return &FileServer{http.FileServer(dir), cache}
}

func (s *FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !s.cache {
		headers := w.Header()
		headers.Set("Cache-Control", "no-cache, no-store, must-revalidate")
		headers.Set("Pragma", "no-cache")
		headers.Set("Expires", "0")
	}
	s.handler.ServeHTTP(w, r)
}
