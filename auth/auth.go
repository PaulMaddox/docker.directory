package auth

import (
	"errors"

	"log"
	"github.com/PaulMaddox/docker.directory/models"
	"github.com/gocraft/web"
	"labix.org/v2/mgo"
)

var (

	// ErrNoDatabase is thrown if no database connection is present
	// during an authentication.
	ErrNoDatabase = errors.New("no database connection")

	// ErrAuthenticationRequired is thrown if no valid user authenticates
	// but authentication is required for the requested resource.
	ErrAuthenticationRequired = errors.New("authentication required")

	// ErrForbidden is thrown if a user is authenticated, but not
	// allowed to access the requested resource.
	ErrForbidden = errors.New("forbidden")

	// ErrInternalServerError is thrown if an internal problem occurs
	// during the authentication process at no fault of the requestor.
	ErrInternalServerError = errors.New("internal server error")
)

// Authenticate is responsible for authenticating users against repositories.
// It current recognises both basic and token authentication (as used by the
// Docker client). It will either return nil,nil if no authentication is required
// for the page, or a user object if authentication succeeds, or an error if
// authentication is required but fails.
func Authenticate(r *web.Request, db *mgo.Session) (*models.User, error) {

	// First check if the URL is whitelisted
	for _, w := range Whitelist {
		// If the method matches...
		if r.Method == w.Method {
			// And the URL matches the regex
			if w.URL.MatchString(r.URL.Path) {
				// This URL is whitelisted
				return nil, nil
			}
		}
	}

	// Ok, so the URL isn't whitelisted; we must authenticate somehow
	// Currently this must be either basic auth or token based.

	// First try to authenticate with basic auth
	user, err := AuthenticateBasic(r, db)
	if err == ErrForbidden {
		// The request attempted to authenticate, but the credentials
		// were incorrect.
		return nil, err
	}

	// From here, there are the following possible scenarios:

	// 1) Basic auth was successful, and the client *isn't* requesting a token to be generated
	if err == nil && r.Header.Get("X-Docker-Token") != "true" {
		log.Printf("Basic auth successful for %s", r.URL.Path)
		return user, nil
	}

	// 2) Basic auth was successful, but the user has requested a token be generated
	// 3) Basic auth failed, but a valid token is present
	// 4) Basic auth failed and no valid token is present
	user, err = AuthenticateToken(r, db, user)
	if err != nil {
		// No valid basic auth or token present
		log.Printf("Failed authentication attempt from %s", r.RemoteAddr)
		return nil, ErrAuthenticationRequired
	}

	return user, nil

}
