package auth

import (
	"encoding/base64"
	"log"
	"strings"
	"labix.org/v2/mgo"

	"github.com/PaulMaddox/docker.directory/models"
	"github.com/gocraft/web"
)

// AuthenticateBasic attempts to authenticate users via basic auth
func AuthenticateBasic(r *web.Request, db *mgo.Session) (*models.User, error) {

	if db == nil {
		return nil, ErrNoDatabase
	}

	auth := r.Header.Get("Authorization")

	// If the request has no authentication credentials, demand them
	if auth == "" {
		return nil, ErrAuthenticationRequired
	}

	parts := strings.Split(auth, " ")
	if len(parts) != 2 || parts[0] != "Basic" {
		return nil, ErrAuthenticationRequired
	}

	decrypted, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, ErrAuthenticationRequired
	}

	creds := strings.Split(string(decrypted[:]), ":")
	if len(creds) != 2 {
		return nil, ErrAuthenticationRequired
	}

	user, err := models.AuthenticateUser(db, creds[0], creds[1])
	if err != nil {
		log.Printf("Invalid login attempt from %s for user %s", r.RemoteAddr, creds[0])
		return nil, ErrAuthenticationRequired
	}

	return user, nil

}
