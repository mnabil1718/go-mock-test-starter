package service

import (
	"testing"

	"github.com/mnabil1718/go-mock-test-starter/entity"
	"github.com/mnabil1718/go-mock-test-starter/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "0").Return(nil)
	category, err := categoryService.Get("0")
	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestCategoryService_GetFound(t *testing.T) {
	var mockCategory = entity.Category{Id: "2", Name: "Mock Category"}
	categoryRepository.Mock.On("FindById", "2").Return(mockCategory)
	category, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, mockCategory.Id, category.Id)
	assert.Equal(t, mockCategory.Name, category.Name)
}
