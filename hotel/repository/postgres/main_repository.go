package postgres

import (
	"database/sql"
	"github.com/doug-martin/goqu"
	"github.com/shellrean-playground/hotel-be-hotel-manager-service/domain"
)

type mainRepository struct {
	tableName string
	db        *goqu.Database
}

func New(db *sql.DB) domain.HotelRepository {
	return &mainRepository{db: goqu.New("default", db), tableName: "hotels"}
}
