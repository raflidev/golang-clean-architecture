package postgres

import (
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func NewPostgreeDB() (*sqlx.DB, error) {
	connstr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"postgres",
		"postgres",
		"golang-clean",
	)

	db, err := sqlx.Connect("pgx", connstr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
