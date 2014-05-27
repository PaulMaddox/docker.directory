package auth

import (
	"fmt"
	"log"
	"regexp"
	"labix.org/v2/mgo"

	"github.com/PaulMaddox/docker.directory/models"
	"github.com/gocraft/web"
)

// AuthenticateToken allows token based access to repositories rather than basic auth
// as per the Docker Registry & Index Spec
func AuthenticateToken(r *web.Request, db *mgo.Session, user *models.User) (*models.User, error) {

	auth := r.Header.Get("Authorization")
	path := r.PathParams["owner"] + "/" + r.PathParams["repository"]

	pattern := fmt.Sprintf(`^Token signature=(%s),repository="(%s)\/(%s)",access=(.*)$`,
		models.TokenPattern,
		models.OwnerPattern,
		models.RepositoryPattern,
	)

	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllStringSubmatch(auth, -1)

	// If we have a successful basic authentication and the docker client
	// has requested that a token be generated then create one, but only if
	// the user can access the requested repository
	if user != nil && len(matches) == 0 {
		if r.Header.Get("X-Docker-Token") == "true" {

			if err := user.CanAccessRepository(db, path); err != nil {
				// The user is not allowed to access this repository location
				log.Printf("Error: User %s is not allowed to access repository %s", user, path)
				return nil, ErrForbidden
			}

			// Attempt to reuse an existing token
			token, err := user.GetAccessToken(db, path)
			if err == nil && token != nil {
				return user, nil
			}

			// Otherwise create a new one
			token, err = models.NewToken(user, path, db)
			if err != nil {
				log.Printf("Error: Unable to create access token (%s)", err)
				return nil, ErrInternalServerError
			}

			log.Printf("Created access token %s for %s to %s", token.Signature, user, token.Path)
			return user, nil

		}
	}

	// Basic auth has failed, but we may have a token we can authenticate with
	if user == nil && len(matches) > 0 {

		sig := matches[0][1]
		token, err := models.LoadTokenBySignature(db, sig)
		if err != nil {
			log.Printf("Error: %s attempted to authenticate with an invalid token (%s)", r.RemoteAddr, err)
			return nil, ErrAuthenticationRequired
		}

		user, err := models.LoadUser(db, token.User)
		if err != nil {
			log.Printf("Error: Cannot load user %s for token %s", token.User, token.Signature)
			return nil, ErrAuthenticationRequired
		}

		log.Printf("Successful token authentication for %s", user)
		return user, nil

	}

	return nil, ErrAuthenticationRequired

}
