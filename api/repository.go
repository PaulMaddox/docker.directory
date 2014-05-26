package api

import (
	"fmt"
	"net/http"

	"github.com/gocraft/web"
)

func (c *APIContext) RepositoryAuth(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}

// RepositoryPut creates a repository if it doesn't already exist
// and if the user has permission to do so.
// PUT http://docker.directory/v1/:namespace/:repository
func (c *APIContext) RepositoryPut(res web.ResponseWriter, req *web.Request) {

}

func (c *APIContext) RepositoryDelete(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}

func (c *APIContext) RepositoryImageGet(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}

func (c *APIContext) RepositoryImagePut(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}
