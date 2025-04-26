package usecase

import (
	"github.com/raflidev/golang-clean-architecture/internal/event"
	"github.com/raflidev/golang-clean-architecture/internal/models"
)

type EventUC struct {
	repo event.Repository
}

func NewEventUseCase(repo event.Repository) event.EventUC {
	return &EventUC{
		repo: repo,
	}
}

func (u *EventUC) FindById(id int64) (*models.Event, error) {
	event, err := u.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	return event, nil
}
