package event

import "github.com/raflidev/golang-clean-architecture/internal/models"

type Repository interface {
	FindById(id int64) (*models.Event, error)
}
