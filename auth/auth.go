package auth

import (
	"errors"
	"regexp"

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

// Authenticate is responsible for authenticating users.
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
		log.Printf("Basic auth successful for %s", user)
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

// Authorize checks whether an authenticated user is allowed to
// access the requested resource. Returns nil if allowed or an error if not
func Authorize(session *mgo.Session, user *models.User, request *web.Request) error {

	// First check if the URL is whitelisted
	for _, w := range Whitelist {
		// If the method matches...
		if request.Method == w.Method {
			// And the URL matches the regex
			if w.URL.MatchString(request.URL.Path) {
				// This URL is whitelisted
				return nil
			}
		}
	}

	if user == nil {
		log.Printf("Attempting to authorize access to %s but no user authenticated", request.URL.Path)
		return ErrAuthenticationRequired
	}

	// Next see what kind of authorization we need.

	// 1) Repository based; see if this user (or any of their organisations)
	//    can access the resource.

	// 2) Image based; see if the image belongs to a repository that
	//    the user (or any of their organisations) has access to.

	imgBased := regexp.MustCompile(`^\/v1/images`)
	repoBased := regexp.MustCompile(`^\/v1\/repositories`)

	switch {
	case imgBased.MatchString(request.URL.Path):

		// TODO: We should restrict access to images based on the parent repository
		// But I can't think of an efficient way of doing this, without searching every single
		// repository, and comparing lists of it's images.
		return nil

	case repoBased.MatchString(request.URL.Path):

		owner := request.PathParams["owner"]

		if owner == user.Username {
			return nil
		}

		orgs, err := user.GetOrganisations(session)
		if err != nil {
			log.Printf("Unable to get organisations for user %s (%s)", user, err)
		}

		for _, org := range *orgs {
			if org.Path == owner {
				return nil
			}
		}

	}

	return ErrForbidden

}
