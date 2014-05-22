package middleware

import "github.com/gocraft/web"

// ContentJson sets the content header for JSON responses
func ContentTypeJson(res web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	res.Header().Set("Content-Type", "application/json")
	next(res, req)
}

// Version adds a X-Docker-Registry-Version cookie header
func Version(res web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	res.Header().Set("X-Docker-Registry-Version", "0.0.1")
	next(res, req)
}
