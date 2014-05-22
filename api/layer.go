package api

import (
	"fmt"
	"net/http"

	"github.com/gocraft/web"
)

func (c *ApiContext) LayerGet(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}

func (c *ApiContext) LayerPut(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}

func (c *ApiContext) LayerDelete(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}
