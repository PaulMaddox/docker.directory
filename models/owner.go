package models

import (
	"fmt"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// Owner is somebody who can access a repository.
// Either a user, or maybe in the future an organisation.
type Owner interface {
	GetID() bson.ObjectId
	CanAccessRepository(session *mgo.Session, path string) error
	fmt.Stringer
}

// OwnerPattern is a regexp for matching and validating repository owners
var OwnerPattern = `[a-zA-Z0-9\._]{4,30}`
