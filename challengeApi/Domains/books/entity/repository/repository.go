package repository

import (
	"FullStack-Challenge/challengeApi/model"
)

type RepositoryInterface interface {
	Create(book *model.Book) error
	GetOne(queryFilter *model.QueryFilters) (model.Book, error)
	GetAll(filter model.QueryFilters) ([]model.Book, error)
	Update(filter *model.QueryFilters, book model.Book) error
	Delete(queryFilter *model.QueryFilters) error
}
