package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocraft/web"
)

func (c *APIContext) ImageIndex(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}

func (c *APIContext) ImageGet(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}

// HTTP 200 if successful
// HTTP 204 if already exists
func (c *APIContext) ImagePut(res web.ResponseWriter, req *web.Request) {
	//res.WriteHeader(http.StatusNotImplemented)
	//fmt.Fprint(res, JSON{"error": "Not implemented"})
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, []byte("[]"))
	log.Printf("Created image %s", req.PathParams["image_id"])
}

func (c *APIContext) ImageDelete(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}
