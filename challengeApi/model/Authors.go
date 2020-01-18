package model

import (
	presenters "FullStack-Challenge/challengeApi/presenters"

	"gopkg.in/mgo.v2/bson"
)

const DBCOLLECTION_AUTHORS = "authors"

type Authors struct {
	Authors []Author `json:"authors,omitempty," bson:"authors,omitempty"`
}

type Author struct {
	Id        bson.ObjectId       `json:"_id,omitempty," bson:"_id,omitempty"`
	FirstName string              `json:"firstName,omitempty," bson:"firstName,omitempty"`
	LastName  string              `json:"lastName,omitempty," bson:"lastName,omitempty"`
	Instance  presenters.Instance `json:"instance,omitempty," bson:"instance,omitempty"`
}
