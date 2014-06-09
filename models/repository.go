package models

import (
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// Repository is a docker repository. It is owned by either a user or an organisation.
type Repository struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Created  time.Time
	Modified time.Time
	Public   bool
	Name     string
	Owner    bson.ObjectId `bson:"owner_id"`
	Images   []string
	Tags     map[string]string
}

// RepositoryPattern is a regexp that can be used for
// matching and validating repository names.
var RepositoryPattern = `[a-zA-Z0-9\._]{4,30}`

// NewRepository creates a new repository object
func NewRepository(owner Owner, name string) *Repository {
	return &Repository{
		Created:  time.Now(),
		Modified: time.Now(),
		Name:     name,
		Public:   false,
		Owner:    owner.GetID(),
	}
}

// LoadRepository loads a repository based on the provided owner
// and repository name
func LoadRepository(session *mgo.Session, owner Owner, name string) (*Repository, error) {

	db := session.Copy()
	defer db.Close()

	collection := db.DB("directory").C("repositories")
	result := &Repository{}
	err := collection.Find(bson.M{
		"owner_id": owner.GetID(),
		"name":     name,
	}).One(result)

	return result, err

}

// Exists checks if a repository already exists
func (r *Repository) Exists(session *mgo.Session) bool {

	db := session.Copy()
	defer db.Close()

	collection := db.DB("directory").C("repositories")

	result := &Repository{}
	err := collection.Find(bson.M{
		"owner_id": r.Owner,
		"name":     r.Name,
	}).One(result)

	return err != mgo.ErrNotFound

}

// Update is a semantic shortcut to Create() as that just runs
// an upsert operation anyway
func (r *Repository) Update(session *mgo.Session) error {
	return r.Create(session)
}

// Create saves a repository to the database
func (r *Repository) Create(session *mgo.Session) error {

	db := session.Copy()
	defer db.Close()

	collection := db.DB("directory").C("repositories")
	_, err := collection.Upsert(bson.M{
		"owner_id": r.Owner,
		"name":     r.Name,
	}, r)

	return err

}

// String provides a textual representation of the repository
func (r *Repository) String() string {
	return r.Owner.String() + "/" + r.Name
}
