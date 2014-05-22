package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PaulMaddox/docker.directory/models"
	"github.com/gocraft/web"
)

func (c *ApiContext) UserLogin(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}

func (c *ApiContext) UserCreate(res web.ResponseWriter, req *web.Request) {

	decoder := json.NewDecoder(req.Body)
	var user models.User
	err := decoder.Decode(&user)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(res, "Invalid JSON")
		return
	}

	if err := user.Create(c.Database); err != nil {
		log.Printf("Failed to create user %s <%s>", user.Username, user.Email)
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(res, err)
		return
	}

	res.WriteHeader(http.StatusCreated)
	log.Printf("Account created successfully for %s <%s>", user.Username, user.Email)
	fmt.Fprintf(res, "User created successfully")

}

func (c *ApiContext) UserUpdate(res web.ResponseWriter, req *web.Request) {
	res.WriteHeader(http.StatusNotImplemented)
	fmt.Fprint(res, Json{"error": "Not implemented"})
}
