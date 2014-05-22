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

	if err := user.Validate(); err != nil {
		log.Printf("Warning: Validation failed for new user %s <%s> (%s)", user.Username, user.Email, err)
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(res, err)
		return
	}

	// The docker tool delivers passwords in plain-text.
	// We have no need for that - bcrypt them
	if err := user.EncryptPassword(); err != nil {
		log.Printf("Error: Unable to bcrypt password for user %s (%s)", user.Username, err)
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, "Could not bcrypt passphrase - aborting!")
	}

	if c.Database == nil {
		log.Printf("Database is nil!")
	}

	db := c.Database.Copy()
	defer db.Close()

	collection := db.DB("directory").C("users")
	if err := collection.Insert(user); err != nil {
		log.Printf("Error: Could not insert new user %s (%s)", user.Username, err)
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, "There was an error while trying to create the user - please try again later")
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
