package postgres

import (
	"context"
	"github.com/doug-martin/goqu"
	"github.com/google/uuid"
	"github.com/shellrean-playground/hotel-be-hotel-manager-service/domain"
	"time"
)

func (m mainRepository) Save(ctx context.Context, hotel *domain.Hotel) (err error) {
	hotel.ID = uuid.New().String()
	hotel.IsDeleted = 0
	hotel.CreatedAt = time.Now().Unix()
	hotel.UpdatedAt = time.Now().Unix()

	crudExec := m.db.From(m.tableName).Insert(goqu.Record{
		"id":         hotel.ID,
		"code":       hotel.Code,
		"name":       hotel.Name,
		"address":    hotel.Address,
		"rate":       hotel.Rate,
		"is_deleted": hotel.IsDeleted,
		"created_at": hotel.CreatedAt,
		"updated_at": hotel.UpdatedAt,
	})
	_, err = crudExec.ExecContext(ctx)
	return
}
