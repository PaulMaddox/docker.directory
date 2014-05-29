package api

import (
	"net/http"

	"github.com/gocraft/web"
)

// GET /images/:image_id/checksum
func (c *Context) ImageChecksumGet(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	res.Write([]byte("Not implemented"))
	return
}

// PUT /images/:image_id/checksum
func (c *Context) ImageChecksumPut(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Not implemented"))
	return
}

// DELETE /images/:image_id/checksum
func (c *Context) ImageChecksumDelete(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	res.Write([]byte("Not implemented"))
	return
}
