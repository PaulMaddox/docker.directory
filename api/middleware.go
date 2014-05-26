package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PaulMaddox/docker.directory/models"
	"github.com/gocraft/web"
)

// The realm to use for all basic authentication
var realm = "Authentication Required"

// Version adds a X-Docker-Registry-Version cookie header
func (c *Context) Version(res web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	res.Header().Set("X-Docker-Registry-Version", "0.0.1")
	next(res, req)
}

// MongoDatabase sets up a database session in the context
func (c *Context) MongoDatabase(res web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.Database = Database
	next(res, req)
}

// ContentTypeJSON sets the content header for JSON responses
func ContentTypeJSON(res web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	res.Header().Set("Content-Type", "application/json")
	next(res, req)
}

// RequestLogger prints a pretty JSON representation of incoming requests
func (c *Context) RequestLogger(res web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {

	b, err := json.MarshalIndent(req, "", "    ")
	if err != nil {
		log.Printf("Error: Unable to create JSON representation of incoming request")
		next(res, req)
		return
	}

	fmt.Print(string(b))
	next(res, req)

}

// AuthenticationWhitelist generates a list of whitelisted URLs
// that require no authentication to access
func (c *Context) AuthenticationWhitelist(res web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {

	c.AuthWhitelist = []string{
		"/v1/users/",
		"/v1/_ping",
	}

	next(res, req)
}

// BasicAuthentication attempts to authenticate users and refuses access
// if authentication fails and the URL is not whitelisted.
func (c *Context) BasicAuthentication(res web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {

	for _, url := range c.AuthWhitelist {
		if req.URL.Path == url {
			// This URL is whitelisted - don't protect it
			next(res, req)
			return
		}
	}

	if c.Database == nil {
		log.Printf("Error: Unable to authenticate user as database connection == nil")
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	auth := req.Header.Get("Authorization")

	// If the request has no authentication credentials, demand them
	if auth == "" {
		res.Header().Set("Www-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
		res.WriteHeader(http.StatusUnauthorized)
		res.Write([]byte("Unauthorized"))
		return
	}

	parts := strings.Split(auth, " ")
	if len(parts) != 2 || parts[0] != "Basic" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Unauthorized"))
		return
	}

	decrypted, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Unauthorized"))
		return
	}

	creds := strings.Split(string(decrypted[:]), ":")
	if len(creds) != 2 {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Unauthorized"))
		return
	}

	user, err := models.AuthenticateUser(c.Database, creds[0], creds[1])
	if err != nil {
		res.Header().Set("Www-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
		res.WriteHeader(http.StatusUnauthorized)
		res.Write([]byte("Unauthorized"))
		return

	}

	// Set the user on the context so that it's available
	// for other modules/handlers
	c.User = user

	next(res, req)

}

// RepositoryToken middleware is responsible for authenticating requests for repositories
// and generating/verifying an access token that is used by the docker client.
func (c *Context) RepositoryToken(res web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {

	if req.Header.Get("X-Docker-Token") == "true" {

		// Ok, the client has asked for a token to be generated.
		// By now, we should have a user authenticated via basic
		// authentication. Throw toys out of pram if not.
		if c.User == nil {
			res.Header().Set("Www-Authenticate", fmt.Sprintf(`Basic real="%s"`, realm))
			res.WriteHeader(http.StatusUnauthorized)
			res.Write([]byte("Unauthorized"))
			return
		}

		// A token request should have a namespace and repository.
		// If not, we have nothing to authenticate the user against.
		if req.PathParams["namespace"] == "" || req.PathParams["repository"] == "" {
			log.Printf("Error: Got an authentication token request from %s but no namespace/repository", c.User)
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("Unauthorized"))
			return
		}

		repository := &models.Repository{
			Namespace: req.PathParams["namespace"],
			Name:      req.PathParams["repository"],
		}

		token, err := repository.AuthenticateOwner(c.User)
		if err != nil {
			log.Printf("Error: User %s is not allowed to access %s", c.User, repository)
			res.WriteHeader(http.StatusForbidden)
			res.Write([]byte("Forbidden"))
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Header().Set("X-Docker-Token", token)
		return

	}

}
