package api

import (
	"fmt"
	"net/http"

	"github.com/gocraft/web"
)

// AncestryGet handles GET requests for retrieving image ancestry information
func (c *ApiContext) AncestryGet(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}

// AncestryPut handles PUT requests for creating image ancestry information
func (c *ApiContext) AncestryPut(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}

// AncestryDelete handles DELETE requests for removing image ancestry information
func (c *ApiContext) AncestryDelete(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}
