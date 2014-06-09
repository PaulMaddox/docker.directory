package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PaulMaddox/docker.directory/storage"
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
	provider := c.User.GetStorageProvider()

	checksum, err := storage.Put(id, req.Body, provider)
	if err != nil {
		log.Printf("Error uploading image %s (%s)", id, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("SHA256: %s", checksum)

	res.WriteHeader(http.StatusOK)
	return

}

// DELETE /images/:image_id/layer
func (c *APIContext) LayerDelete(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}
