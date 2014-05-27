package api

import (
	"net/http"

	"github.com/gocraft/web"
)

func (c *Context) ImageChecksumGet(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	res.Write([]byte("Not implemented"))
	return
}

func (c *Context) ImageChecksumPut(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Not implemented"))
	return
}

func (c *Context) ImageChecksumDelete(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	res.Write([]byte("Not implemented"))
	return
}
