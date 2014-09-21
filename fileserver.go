package main

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
	if cache {
		http.Handle("/", http.FileServer(http.Dir(dir)))
	} else {
		http.Handle("/", disableCache(http.FileServer(http.Dir(dir))))
	}
	return http.ListenAndServe(address, nil)
}

func disableCache(h http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		h.ServeHTTP(rw, r)
	})
}
