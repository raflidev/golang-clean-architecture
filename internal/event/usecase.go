package event

import "github.com/raflidev/golang-clean-architecture/internal/models"

type EventUC interface {
	FindById(id int64) (*models.Event, error)
}
