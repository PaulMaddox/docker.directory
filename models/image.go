package models

import (
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// Image stores a Docker image and it's checksum
type Image struct {
	ID              bson.ObjectId `bson:"_id,omitempty"`
	Identifier      string        `json:"id"`
	Checksum        string
	Author          string
	Parent          string
	Created         time.Time
	Container       string
	DockerVersion   string
	Tag             string
	Architecture    string
	OperatingSystem string
	Size            int64
	Config          interface{} `json:"config"`
	ContainerConfig interface{} `json:"container_config"`
}

// LoadImage loads an image from the database based on
// the provided identifier. Returns the loaded image if successful
// otherwise returns an error.
func LoadImage(session *mgo.Session, identifier string) (*Image, error) {

	db := session.Copy()
	defer db.Close()

	collection := db.DB("directory").C("images")

	result := &Image{}
	err := collection.Find(bson.M{
		"identifier": identifier,
	}).One(result)

	return result, err

}

// Create saves the image to the database
func (i *Image) Create(session *mgo.Session) error {

	db := session.Copy()
	defer db.Close()

	collection := db.DB("directory").C("images")
	return collection.Insert(i)

}

// Exists checks if an image already exists in our database
func (i *Image) Exists(session *mgo.Session) bool {

	db := session.Copy()
	defer db.Close()

	collection := db.DB("directory").C("images")

	result := &Image{}
	err := collection.Find(bson.M{
		"identifier": i.Identifier,
	}).One(result)

	return err != mgo.ErrNotFound

}
