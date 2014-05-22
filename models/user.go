package models

import (
	"errors"
	"log"
	"regexp"

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
func (u *User) Validate() error {

	// Check the username
	pattern := `^[a-z0-9]{4,30}$`
	if r, err := regexp.Compile(pattern); err != nil {
		log.Printf("Warning: Invalid regular expression for verifying user")
	} else {
		if !r.MatchString(u.Username) {
			return errors.New("Wrong username format (should match " + pattern + ")")
		}
	}

	// Check if the username/email already exists
	// TODO: Implement this
	if false {
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
