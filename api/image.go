package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"labix.org/v2/mgo"

	"github.com/PaulMaddox/docker.directory/models"
	"github.com/gocraft/web"
)

// GET /images
func (c *APIContext) ImageIndex(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}

// ImageGet loads an image from the database and returns it to the requestor
// in JSON format.
// GET /images/:image_id/json
func (c *APIContext) ImageGet(res web.ResponseWriter, req *web.Request) {

	id := req.PathParams["image_id"]
	img, err := models.LoadImage(c.Database, req.PathParams["image_id"])
	if err != nil {
		if err == mgo.ErrNotFound {
			res.WriteHeader(http.StatusNotFound)
			return
		} else {
			log.Printf("Error while loading image %s from database for user %s (%s)", id, c.User, err)
			res.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// Write it out as json
	data, err := json.Marshal(img)
	if err != nil {
		log.Printf("Error while marshalling image %s to JSON for user %s (%s)", id, c.User, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, string(data[:]))

}

// ImagePut recieves docker image metadata in json form and stores it
// in the database if it doesn't already exist.
// PUT /images/:image_id/json
func (c *APIContext) ImagePut(res web.ResponseWriter, req *web.Request) {

	id := req.PathParams["image_id"]

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Invalid data for image %s from user %s", id, c.User)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	img := &models.Image{}
	if err := json.Unmarshal(data, img); err != nil {
		log.Printf("Invalid JSON for image %s from user %s (%s)", id, c.User, err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	if img.Exists(c.Database) {
		log.Printf("Ignoring previously uploaded image %s", id)
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(res, JSON{"error": "Image already exists"})
		return
	}

	if err := img.Create(c.Database); err != nil {
		log.Printf("Error while saving image %s for user %s (%s)", id, c.User, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
	log.Printf("Created image %s", req.PathParams["image_id"])

	return

}

// DELETE /images/:image_id/json
func (c *APIContext) ImageDelete(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}
