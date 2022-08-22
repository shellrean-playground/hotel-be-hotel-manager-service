package postgres

import (
	"context"
	"github.com/doug-martin/goqu"
	"github.com/shellrean-playground/hotel-be-hotel-manager-service/domain"
	"time"
)

func (m mainRepository) Update(ctx context.Context, hotel *domain.Hotel) (err error) {
	hotel.UpdatedAt = time.Now().Unix()

	crudExec := m.db.From(m.tableName).
		Where(goqu.Ex{"id": hotel.ID}).
		Insert(goqu.Record{
			"code":       hotel.Code,
			"name":       hotel.Name,
			"address":    hotel.Address,
			"rate":       hotel.Rate,
			"updated_at": hotel.UpdatedAt,
		})
	_, err = crudExec.ExecContext(ctx)
	return
}
