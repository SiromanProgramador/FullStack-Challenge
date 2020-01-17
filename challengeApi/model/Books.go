package model

import (
	presenters "FullStack-Challenge/challengeApi/presenters"

	"gopkg.in/mgo.v2/bson"
)

const DBCOLLECTION_BOOKS = "books"

type Books struct {
	Books []Book `json:"books,omitempty," bson:"books,omitempty"`
}

type Book struct {
	Id       bson.ObjectId       `json:"_id,omitempty," bson:"_id,omitempty"`
	Name     string              `json:"name,omitempty," bson:"name,omitempty"`
	ISBN     string              `json:"isbn,omitempty," bson:"isbn,omitempty"`
	AuthorId bson.ObjectId       `json:"authorId,omitempty," bson:"authorId,omitempty"`
	Instance presenters.Instance `json:"instance,omitempty," bson:"instance,omitempty"`
}
