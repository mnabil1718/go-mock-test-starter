package repository

import "github.com/mnabil1718/go-mock-test-starter/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
