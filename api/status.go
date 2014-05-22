package api

import (
	"fmt"

	"github.com/gocraft/web"
)

// Status reports on the status of the registry.
// This endpoint is also used to determin if the registry supports SSL
func (c *ApiContext) Status(res web.ResponseWriter, req *web.Request) {
	fmt.Fprintf(res, "Not implemented")
}
