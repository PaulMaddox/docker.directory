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
	fmt.Stringer
}

// OwnerPattern is a regexp for matching and validating repository owners
var OwnerPattern = `[a-zA-Z0-9\._]{4,30}`

// LoadOwner loads an owner from either a username or organisation name
func LoadOwnerByName(session *mgo.Session, name string) (Owner, error) {

	db := session.Copy()
	defer db.Close()

	// First try and look for a user with this name
	users := db.DB("directory").C("users")
	user := &User{}
	err := users.Find(bson.M{
		"username": name,
	}).One(user)

	if err == nil {
		return user, nil
	}

	// Otherwise look for an organisation with this name
	organisations := db.DB("directory").C("organisations")
	organisation := &Organisation{}
	err = organisations.Find(bson.M{
		"path": name,
	}).One(organisation)

	if err == nil {
		return organisation, nil
	}

	return nil, mgo.ErrNotFound

}

// LoadOwner loads an owner from either a username or organisation name
func LoadOwnerByID(session *mgo.Session, id bson.ObjectId) (Owner, error) {

	db := session.Copy()
	defer db.Close()

	// First try and look for a user with this name
	users := db.DB("directory").C("users")
	user := &User{}
	err := users.Find(bson.M{
		"_id": id,
	}).One(user)

	if err == nil {
		return user, nil
	}

	// Otherwise look for an organisation with this name
	organisations := db.DB("directory").C("organisations")
	organisation := &Organisation{}
	err = organisations.Find(bson.M{
		"_id": id,
	}).One(organisation)

	if err == nil {
		return organisation, nil
	}

	return nil, mgo.ErrNotFound

}
