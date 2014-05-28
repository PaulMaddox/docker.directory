package api

import (
	"fmt"
	"net/http"

	"github.com/gocraft/web"
)

// Status reports on the status of the registry.
// This endpoint is also used to determin if the registry supports SSL
func (c *APIContext) Status(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("X-Docker-Registry-Version", "0.0.1")
	res.Header().Set("X-Docker-Registry-Standalone", "false")
	fmt.Fprint(res, JSON{"status": "ok"})
}
