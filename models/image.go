package models

import "labix.org/v2/mgo/bson"

// Image stores a Docker image and it's checksum
type Image struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Identifier string        `json:"id"`
	Tag        string
	Checksum   string
	Metadata   string
}
