package api

import (
	"fmt"
	"net/http"

	"github.com/gocraft/web"
)

// GET /repositories/:owner/:repository/tags
func (c *APIContext) TagsGet(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}

// GET /repositories/:owner/:repository/tags/:tag
func (c *APIContext) TagGet(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}

// PUT /repositories/:owner/:repository/tags/:tag
func (c *APIContext) TagPut(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}

// DELETE /repositories/:owner/:repository/tags/:tag
func (c *APIContext) TagDelete(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}
