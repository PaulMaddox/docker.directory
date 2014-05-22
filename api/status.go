package api

import (
	"fmt"
	"net/http"

	"github.com/gocraft/web"
)

// Status reports on the status of the registry.
// This endpoint is also used to determin if the registry supports SSL
func (c *ApiContext) Status(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}
