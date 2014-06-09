package models

import "labix.org/v2/mgo/bson"

type Organisation struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Path string
	Name string
}

func (o *Organisation) GetID() bson.ObjectId {
	return o.ID
}

func (o *Organisation) String() string {
	return o.Name
}
