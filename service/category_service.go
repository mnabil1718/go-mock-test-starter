package service

import (
	"errors"

	"github.com/mnabil1718/go-mock-test-starter/entity"
	"github.com/mnabil1718/go-mock-test-starter/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return category, errors.New("category not found")
	} else {
		return category, nil
	}
}
