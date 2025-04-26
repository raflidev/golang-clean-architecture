package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/raflidev/golang-clean-architecture/internal/event"
	"github.com/raflidev/golang-clean-architecture/internal/models"
)

type EventRepository struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) event.Repository {
	return &EventRepository{db: db}
}

func (r *EventRepository) FindById(id int64) (*models.Event, error) {
	stmt := `SELECT * FROM event WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	event := &models.Event{}
	args := []any{id}
	err := r.db.QueryRowxContext(ctx, stmt, args...).StructScan(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
