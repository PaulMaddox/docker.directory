package api

import (
	"fmt"
	"net/http"

	"github.com/gocraft/web"
)

func (c *ApiContext) RepositoryAuth(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}

func (c *ApiContext) RepositoryPut(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}

func (c *ApiContext) RepositoryDelete(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}

func (c *ApiContext) RepositoryImageGet(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}

func (c *ApiContext) RepositoryImagePut(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}
