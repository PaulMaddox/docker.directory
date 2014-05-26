package models

import

// A repository belongs to either a user or an organisation
// and holds a number of Docker images

"github.com/nu7hatch/gouuid"

type Repository struct {
	Id        string
	Created   int64
	Modified  int64
	Name      string
	Namespace string
}

func (r *Repository) Exists() {

}

// Authenticate checks that to see if an owner (typically a user or organisation)
// has sufficient rights for this repository. If successful it wil return an access
// token that can be used for future access to this repository. If not, returns an error.
func (r *Repository) AuthenticateOwner(owner Owner) (string, error) {

	// TODO: Actually authenticate
	token, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return token.String(), nil

}

// String provides a textual representation of the repository
func (r *Repository) String() string {
	return r.Namespace + "/" + r.Name
}
