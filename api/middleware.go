package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gocraft/web"
)

// Version adds a X-Docker-Registry-Version cookie header
func (c *Context) Version(res web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	res.Header().Set("X-Docker-Registry-Version", "0.0.1")
	next(res, req)
}

// MongoDatabase sets up a database session in the context
func (c *Context) MongoDatabase(res web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.Database = Database
	next(res, req)
}

// ContentJson sets the content header for JSON responses
func ContentTypeJson(res web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	res.Header().Set("Content-Type", "application/json")
	next(res, req)
}

// RequestLogger prints a pretty JSON representation of incoming requests
func RequestLogger(res web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {

	b, err := json.MarshalIndent(req, "", "    ")
	if err != nil {
		log.Printf("Unable to create JSON representation of incoming request")
		next(res, req)
		return
	}

	fmt.Print(string(b))
	next(res, req)

}
