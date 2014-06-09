package api

import (
	"log"
	"net/http"
	"labix.org/v2/mgo"

	"github.com/PaulMaddox/docker.directory/models"
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

	// Check the image exists
	id := req.PathParams["image_id"]
	img, err := models.LoadImage(c.Database, id)
	if err != nil {
		if err == mgo.ErrNotFound {
			res.WriteHeader(http.StatusNotFound)
			return
		} else {
			log.Printf("Unable to load image %s for user (%s)", c.User, err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// TODO: Before Docker 0.1.0 X-Docker-Checksum was used here
	// with tarsum validation. Do we need to offer backwards compat?
	checksum := req.Header.Get("X-Docker-Checksum-Payload")
	if checksum == "" {
		log.Printf("Invalid checksum %s for image %s from user %s", checksum, id, c.User)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	img.Checksum = checksum
	if err := img.Update(c.Database); err != nil {
		log.Printf("Failed to save image %s for user %s (%s)", id, c.User, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	return

}

// DELETE /images/:image_id/checksum
func (c *Context) ImageChecksumDelete(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	res.Write([]byte("Not implemented"))
	return
}
