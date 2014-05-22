package api

import (
	"fmt"
	"net/http"

	"github.com/gocraft/web"
)

func (c *ApiContext) UserLogin(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}

func (c *ApiContext) UserCreate(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}

func (c *ApiContext) UserUpdate(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}
