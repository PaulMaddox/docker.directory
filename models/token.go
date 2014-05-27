package models

import (
	"fmt"
	"time"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/nu7hatch/gouuid"
)

// TokenDuration is the default token expiration time
const TokenDuration = 1 * time.Hour

// Token provides access for a user to a repository
type Token struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Created     time.Time
	Modified    time.Time
	User        bson.ObjectId `bson:"user_id"`
	Path        string
	Signature   string
	ReadAccess  bool
	WriteAccess bool
	Expires     time.Time
}

// TokenPattern is a regexp pattern that can be used for matching tokens (UUIDs)
var TokenPattern = `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`

// NewToken creates a new access token for a given user + repository path.
// This will be automatically served by the Authenticate middleware.
func NewToken(user *User, path string, db *mgo.Session) (*Token, error) {

	sig, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	token := &Token{
		Created:     time.Now(),
		Modified:    time.Now(),
		User:        user.ID,
		Path:        path,
		Signature:   sig.String(),
		ReadAccess:  true,
		WriteAccess: true,
		Expires:     time.Now().Add(TokenDuration),
	}

	session := db.Copy()
	defer session.Close()

	if err := session.DB("directory").C("tokens").Insert(token); err != nil {
		return nil, err
	}

	return token, nil

}

// String returns the token ready to be placed in a HTTP header
func (t *Token) String() string {
	return fmt.Sprintf(`signature=%s,repository="%s",access=%s`,
		t.Signature,
		t.Path,
		"readwrite",
	)
}

// LoadTokenBySignature loads a token from the database
func LoadTokenBySignature(session *mgo.Session, signature string) (*Token, error) {

	db := session.Copy()
	defer db.Close()

	collection := db.DB("directory").C("tokens")

	result := &Token{}
	err := collection.Find(bson.M{
		"signature": signature,
	}).One(result)

	return result, err

}
