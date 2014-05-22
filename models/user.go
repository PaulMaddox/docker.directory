package models

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"code.google.com/p/go.crypto/bcrypt"
)

type User struct {
	Email    string
	Username string
	Password string
}

// Validate checks that a user meets the following requirements as
// defined bny the Docker Registry & Index Spec
// * Username minimum of 4 characters
// * Username must match [a-z0-9]{4,30}
// * Username/email not already in use
// * Password between 5-30 characters
func (u *User) Validate(session *mgo.Session) error {

	// Check the username
	pattern := `^[a-zA-Z0-9\._]{4,30}$`
	if r, err := regexp.Compile(pattern); err != nil {
		log.Printf("Warning: Invalid regular expression for verifying user")
	} else {
		if !r.MatchString(u.Username) {
			return errors.New("Wrong username format (should match " + pattern + ")")
		}
	}

	// Check if the username/email already exists
	if u.Exists(session) {
		return errors.New("Username or email already exists")
	}

	// Check the password
	if len(u.Password) < 5 || len(u.Password) > 30 {
		return errors.New("Invalid password (should be between 5-30 characters)")
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

// Create creates a new user in the database via the provided Mongo session
func (u *User) Create(session *mgo.Session) error {

	if err := u.Validate(session); err != nil {
		msg := fmt.Sprintf("Validation failed for new user %s <%s> (%s)", u.Username, u.Email, err)
		log.Print(msg)
		return errors.New(msg)
	}

	// The docker tool delivers passwords in plain-text.
	// We have no need for that - bcrypt them
	if err := u.EncryptPassword(); err != nil {
		msg := fmt.Sprintf("Unable to bcrypt password for user %s (%s)", u.Username, err)
		log.Print(msg)
		return errors.New(msg)
	}

	db := session.Copy()
	defer db.Close()

	collection := db.DB("directory").C("users")
	if err := collection.Insert(u); err != nil {
		msg := fmt.Sprintf("Could not insert new user %s (%s)", u.Username, err)
		log.Printf(msg)
		return errors.New(msg)
	}

	return nil

}
