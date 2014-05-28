package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PaulMaddox/docker.directory/models"
	"github.com/gocraft/web"
)

// UserCreate handles user creation POST requests from the docker client
func (c *APIContext) UserCreate(res web.ResponseWriter, req *web.Request) {

	decoder := json.NewDecoder(req.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		log.Printf("Failed to create user (%s)", err)
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, `"Invalid JSON"`)
		return
	}

	if err := user.Create(c.Database); err != nil {
		log.Printf("Failed to create user %s <%s>", user.Username, user.Email)
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, `"%s"`, err)
		return
	}

	res.WriteHeader(http.StatusCreated)
	log.Printf("Account created successfully for %s <%s>", user.Username, user.Email)
	fmt.Fprintf(res, `"User created successfully"`)

}

func (c *APIContext) UserUpdate(res web.ResponseWriter, req *web.Request) {
	log.Printf("User update")
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, JSON{"error": "Not implemented"})
}

// UserLogin authenticates a user
func (c *APIContext) UserLogin(res web.ResponseWriter, req *web.Request) {
	if c.User != nil {
		res.WriteHeader(http.StatusOK)
		return
	} else {
		res.WriteHeader(http.StatusUnauthorized)
		return
	}
}
