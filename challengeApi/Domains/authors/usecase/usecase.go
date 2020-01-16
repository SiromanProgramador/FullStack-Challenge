package usecase

import (
	"challengeApi/Domains/authors/entity/repository"
	"challengeApi/model"
)

type UsecaseInterface interface {
	Create(author model.Author) (model.Author, error)
	Delete(filfer model.QueryFilters) error
	GetOne(queryFilters model.QueryFilters) (model.Author, error)
	GetAll(filter model.QueryFilters) ([]model.Author, error)
	Update(filter *model.QueryFilters, objectToUpdate model.Author) error
}

type Usecase struct {
	repo repository.RepositoryInterface
}

func NewUsecase(repo repository.RepositoryInterface) UsecaseInterface {
	return &Usecase{
		repo: repo,
	}
}
func (u *Usecase) Create(objectToCreate model.Author) (model.Author, error) {

	errCreate := u.repo.Create(&objectToCreate)
	return objectToCreate, errCreate
}

func (u *Usecase) Delete(filter model.QueryFilters) error {

	return u.repo.Delete(&filter)
}

func (u *Usecase) GetOne(queryFilters model.QueryFilters) (model.Author, error) {

	return u.repo.GetOne(&queryFilters)
}

func (u *Usecase) GetAll(filter model.QueryFilters) ([]model.Author, error) {

	return u.repo.GetAll(filter)
}

func (u *Usecase) Update(filter *model.QueryFilters, objectToUpdate model.Author) error {

	err := u.repo.Update(filter, objectToUpdate)
	return err
}
