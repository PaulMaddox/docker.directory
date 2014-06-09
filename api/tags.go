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

// GET /repositories/:owner/:repository/tags
func (c *APIContext) TagsGet(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}

// GET /repositories/:owner/:repository/tags/:tag
func (c *APIContext) TagGet(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}

// PUT /repositories/:owner/:repository/tags/:tag
func (c *APIContext) TagPut(res web.ResponseWriter, req *web.Request) {

	owner, err := models.LoadOwnerByName(c.Database, req.PathParams["owner"])
	if err != nil {
		log.Printf("Unable to load repositories for owner %s (%s)", req.PathParams["owner"], err)
		res.WriteHeader(http.StatusNotFound)
		return
	}

	repository, err := models.LoadRepository(c.Database, owner, req.PathParams["repository"])
	if err != nil {
		log.Printf("Unable to load repository %s (%s)", req.PathParams["repository"], err)
		res.WriteHeader(http.StatusNotFound)
		return
	}

	// Read the request body
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Invalid JSON for tag %s/%s:%s %s (%s)", req.PathParams["owner"], req.PathParams["repository"], req.PathParams["tag"], err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	var image string
	if err := json.Unmarshal(data, &image); err != nil {
		log.Printf("Invalid JSON for tag %s/%s:%s %s (%s)", req.PathParams["owner"], req.PathParams["repository"], req.PathParams["tag"], err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	repository.Tags[req.PathParams["tag"]] = image

	if err := repository.Update(c.Database); err != nil {
		log.Printf("Failed to save tag %s to repository %s (%s)", req.PathParams["tag"], repository, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("Saved tag %s to repository %s", req.PathParams["tag"], repository)

}

// DELETE /repositories/:owner/:repository/tags/:tag
func (c *APIContext) TagDelete(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}
