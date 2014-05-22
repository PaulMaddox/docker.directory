package api

import (
	"fmt"
	"net/http"

	"github.com/gocraft/web"
)

func (c *ApiContext) ImageIndex(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}

func (c *ApiContext) ImageGet(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}

func (c *ApiContext) ImagePut(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}

func (c *ApiContext) ImageDelete(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}
