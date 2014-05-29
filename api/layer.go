package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gocraft/web"
)

// GET /images/:image_id/layer
func (c *APIContext) LayerGet(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}

// PUT /images/:image_id/layer
func (c *APIContext) LayerPut(res web.ResponseWriter, req *web.Request) {

	id := req.PathParams["image_id"]
	storage := c.User.GetStorageProvider()

	err := storage.Put(id, req.Body)
	if err != nil {
		log.Printf("Error uploading image %s (%s)", id, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return

}

// DELETE /images/:image_id/layer
func (c *APIContext) LayerDelete(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}
