package repository

import (
	"FullStack-Challenge/challengeApi/model"
)

type RepositoryInterface interface {
	Create(teacher *model.Author) error
	GetOne(queryFilter *model.QueryFilters) (model.Author, error)
	GetAll(filter model.QueryFilters) ([]model.Author, error)
	Update(filter *model.QueryFilters, user model.Author) error
	Delete(queryFilter *model.QueryFilters) error
}
