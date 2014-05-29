package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PaulMaddox/docker.directory/models"
	"github.com/gocraft/web"
)

// PUT /repositories/:owner/:repository/auth
func (c *APIContext) RepositoryAuth(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}

// RepositoryPut creates a repository if it doesn't already exist
// and if the user has permission to do so.
// PUT http://docker.directory/v1/:owner/:repository
func (c *APIContext) RepositoryPut(res web.ResponseWriter, req *web.Request) {

	path := req.PathParams["owner"] + "/" + req.PathParams["repository"]

	// Check if the user is allowed to create this repository
	if err := c.User.CanAccessRepository(c.Database, path); err != nil {
		log.Printf("User %s is not allowed to access %s", c.User, path)
		res.WriteHeader(http.StatusForbidden)
		res.Write([]byte("Forbidden"))
		return
	}

	// TODO: This should be either the organisation or the user
	repository := models.NewRepository(
		c.User,
		req.PathParams["repository"],
	)

	var images []models.Image
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Invalid data for repository %s from user %s", path, c.User)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(data, &images); err != nil {
		log.Printf("Invalid JSON for repository %s from user %s (%s)", path, c.User, err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: I suspect that if the images change, we need to update the record
	repository.Images = images

	if repository.Exists(c.Database) {
		log.Printf("User %s is trying to create repository %s but it already exists", c.User, path)
		res.WriteHeader(http.StatusCreated)
		return
	}

	if err := repository.Create(c.Database); err != nil {
		log.Printf("Error: Unable to save repository %s (%s)", path, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("User %s created repository %s", c.User, path)
	res.WriteHeader(http.StatusCreated)

}

// DELETE /repositories/:owner/:repository
func (c *APIContext) RepositoryDelete(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}

// GET /repositories/:owner/:repository/images
func (c *APIContext) RepositoryImageGet(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}

// PUT /repositories/:owner/:repository/images
func (c *APIContext) RepositoryImagePut(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}
