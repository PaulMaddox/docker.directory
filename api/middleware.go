package api

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PaulMaddox/docker.directory/auth"
	"github.com/gocraft/web"
)

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

	log.Printf("***************************************************************************************")
	log.Printf("* %s: %s", req.Method, req.URL.Path)
	log.Printf("***************************************************************************************")
	for header := range req.Header {
		log.Printf("* %s: %s", header, req.Header[header])
	}

	content := req.Header.Get("Content-Type")
	if content == "text/plain" || content == "application/json" {

		body, _ := ioutil.ReadAll(req.Body)
		if len(body) > 0 {
			log.Printf("* %s", string(body[:]))
		}

		// As we've now drained req.Body, we need
		// to refill it so other middleware/handlers
		// don't get an empty body.
		restore := bytes.NewReader(body)
		req.Body = ioutil.NopCloser(restore)

	}

	log.Printf("***************************************************************************************")

	next(res, req)
}

// Authenticate middleware authenticates a HTTP request either via basic auth or by accesstoken
// as per the Docker Registry & Index Spec.
func (c *Context) Authenticate(w web.ResponseWriter, r *web.Request, next web.NextMiddlewareFunc) {

	user, err := auth.Authenticate(r, c.Database)
	if err != nil {
		switch err {
		case auth.ErrAuthenticationRequired:
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("Www-Authenticate", `Basic realm="Authentication Required"`)
			w.Write([]byte("Unauthorized"))
			return
		case auth.ErrForbidden:
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Forbidden"))
			return
		case auth.ErrNoDatabase:
			w.WriteHeader(http.StatusInternalServerError)
			return
		default:
			return
		}
	}

	if user != nil {

		// Set the user on the context so that it can be used by other
		// middleware and handlers
		c.User = user

		// If the user has a token for this repository, inject it
		path := r.PathParams["owner"] + "/" + r.PathParams["repository"]
		token, err := c.User.GetAccessToken(c.Database, path)
		if err != nil {
			log.Printf("Unable to get token for %s to access %s (%s)", user, path, err)
		}

		if token != nil {
			w.Header().Set("X-Docker-Token", token.String())
			w.Header().Set("X-Docker-Endpoints", r.Host)
		}

	}

	next(w, r)

}
