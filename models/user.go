package models

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"time"
	"github.com/PaulMaddox/docker.directory/storage"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"code.google.com/p/go.crypto/bcrypt"
)

// User encapulates an authenticated user and is used
// by the various contexts/handlers
type User struct {
	ID              bson.ObjectId `bson:"_id,omitempty"`
	Created         time.Time
	Modified        time.Time
	Email           string
	Username        string
	Password        string
	StorageProvider int
	AwsAccessKey    string
	AwsSecretKey    string
	AwsBucket       string
}

// String prints a textual representation of the user
func (u *User) String() string {
	return u.Username + " <" + u.Email + ">"
}

// Validate checks that a user meets the following requirements as
// defined bny the Docker Registry & Index Spec
// * Username minimum of 4 characters
// * Username must match [a-z0-9]{4,30}
// * Username/email not already in use
// * Password between 5-30 characters
func (u *User) Validate(session *mgo.Session) error {

	// Check the username
	if r, err := regexp.Compile(OwnerPattern); err != nil {
		log.Printf("invalid regular expression for verifying user")
	} else {
		if !r.MatchString(u.Username) {
			return errors.New("wrong username format (should match " + OwnerPattern + ")")
		}
	}

	// Check if the username/email already exists
	if u.Exists(session) {
		return errors.New("username or email already exists")
	}

	// Check the password
	if len(u.Password) < 5 || len(u.Password) > 30 {
		return errors.New("invalid password (should be between 5-30 characters)")
	}

	return nil

}

// EncryptPassword converts a plaintext password (as delivered by the 'docker login' program)
// to a bcrypted one, ready for storing in our DB.
func (u *User) EncryptPassword() error {

	p := []byte(u.Password)
	c, err := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(c[:])

	// Overwrite the plain-text password in memory
	for i := 0; i < len(p); i++ {
		p[i] = 0
	}

	return nil

}

// Exists checks if a user exists in the database via the provided Mongo session
func (u *User) Exists(session *mgo.Session) bool {

	db := session.Copy()
	defer db.Close()

	collection := db.DB("directory").C("users")

	result := &User{}
	err := collection.Find(bson.M{
		"email":    u.Email,
		"username": u.Username,
	}).One(result)

	return err != mgo.ErrNotFound

}

// LoadUser fetches a user from the database based on it's ObjectId
func LoadUser(session *mgo.Session, id bson.ObjectId) (*User, error) {

	db := session.Copy()
	defer db.Close()

	collection := db.DB("directory").C("users")

	result := &User{}
	err := collection.Find(bson.M{
		"_id": id,
	}).One(result)

	return result, err

}

// Create creates a new user in the database via the provided Mongo session
func (u *User) Create(session *mgo.Session) error {

	if err := u.Validate(session); err != nil {
		msg := fmt.Sprintf("validation failed for new user %s <%s> (%s)", u.Username, u.Email, err)
		log.Print(msg)
		return errors.New(msg)
	}

	// The docker tool delivers passwords in plain-text.
	// We have no need for that - bcrypt them
	if err := u.EncryptPassword(); err != nil {
		msg := fmt.Sprintf("unable to bcrypt password for user %s (%s)", u.Username, err)
		log.Print(msg)
		return errors.New(msg)
	}

	db := session.Copy()
	defer db.Close()

	// By default, all users use S3 as their storage provider
	// although this could be overridden in the Web IU at a later time
	u.StorageProvider = storage.ProviderAws

	// TODO: Remove this at later stage
	u.StorageProvider = storage.ProviderDummy

	// Set the created/modified dates on the record
	u.Created = time.Now()
	u.Modified = time.Now()

	collection := db.DB("directory").C("users")
	if err := collection.Insert(u); err != nil {
		msg := fmt.Sprintf("could not insert new user %s (%s)", u.Username, err)
		log.Printf(msg)
		return errors.New(msg)
	}

	return nil

}

// AuthenticateUser loads a user from the database session provided, queryinng
// by username and verifies the password. Returns an error if the user is unable to be authenticated.
func AuthenticateUser(session *mgo.Session, username, password string) (*User, error) {

	db := session.Copy()
	defer db.Close()

	collection := db.DB("directory").C("users")

	result := &User{}
	err := collection.Find(bson.M{
		"username": username,
	}).One(result)

	if err != nil {
		msg := fmt.Sprintf("unable to authorize non existant user %s", username)
		return result, errors.New(msg)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		msg := fmt.Sprintf("bad authentication attempt for user %s", username)
		return result, errors.New(msg)
	}

	return result, nil

}

// GetStorageProvider fetches a storage provider configured to the user's preference (eg: AWS S3)
func (u *User) GetStorageProvider() storage.StorageProvider {

	switch u.StorageProvider {
	case storage.ProviderAws:
		return storage.NewAwsProvider(u.AwsAccessKey, u.AwsSecretKey, u.AwsBucket)
	case storage.ProviderDummy:
		return storage.NewDummyProvider()
	}

	return nil

}

// CanAccessRepository checks if a user can access a repository.
// This satisfies the models.Owner interface. Returns nil if they can,
// or an error if not.
func (u *User) CanAccessRepository(session *mgo.Session, path string) error {
	// TODO: Check if a user can access a repository
	return nil
}

// GetAccessToken fetches a previously created token that allows this
// user access to the provided repository path - if it finds a valid
// one in the database that hasn't expired.
func (u *User) GetAccessToken(db *mgo.Session, path string) (*Token, error) {

	session := db.Copy()
	defer session.Close()

	collection := session.DB("directory").C("tokens")

	var tokens []Token
	collection.Find(bson.M{
		"user_id": u.ID,
		"path":    path,
	}).All(&tokens)

	for index := range tokens {
		if time.Now().After(tokens[index].Expires) {
			// Token has expired
			continue
		}
		return &tokens[index], nil
	}

	return nil, nil

}

// GetID is a helper function that satisfies the Owner interface
func (u *User) GetID() bson.ObjectId {
	return u.ID
}
