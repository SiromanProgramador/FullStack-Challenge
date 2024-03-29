package mongodb

import (
	"FullStack-Challenge/challengeApi/Domains/books/entity/repository"
	"FullStack-Challenge/challengeApi/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repo struct {
	session *mgo.Session
}

func NewMongoDBRepository(session *mgo.Session) repository.RepositoryInterface {
	return &repo{
		session: session,
	}
}

func setCollection(session *mgo.Session, colName string) *mgo.Collection {
	collection := session.DB("challengeDB").C(colName)
	return collection
}

const CollectionName = model.DBCOLLECTION_BOOKS

func (r *repo) Create(objectToCreate *model.Book) error {
	collection := setCollection(r.session, CollectionName)
	errInsert := collection.Insert(objectToCreate)

	return errInsert
}

func (r *repo) Delete(queryFilter *model.QueryFilters) error {

	collection := setCollection(r.session, CollectionName)
	errDelete := collection.Remove(queryFilter.Filter)

	return errDelete
}

func (r *repo) GetOne(queryFilter *model.QueryFilters) (objectToReturn model.Book, errFind error) {
	collection := setCollection(r.session, CollectionName)

	errFind = collection.Find(queryFilter.Filter).One(&objectToReturn)

	return objectToReturn, errFind
}

func (r *repo) GetAll(filter model.QueryFilters) (objectsToReturn []model.Book, errFind error) {

	collection := setCollection(r.session, CollectionName)

	errFind = collection.Find(filter.Filter).Select(filter.Select).Sort(filter.Sort...).Limit(filter.Limit).Skip(filter.Skip).All(&objectsToReturn)

	return objectsToReturn, errFind
}

func (r *repo) Update(filter *model.QueryFilters, objectToUpdate model.Book) error {
	collection := setCollection(r.session, CollectionName)

	errUpdate := collection.Update(filter.Filter, bson.M{"$set": objectToUpdate})

	return errUpdate
}
