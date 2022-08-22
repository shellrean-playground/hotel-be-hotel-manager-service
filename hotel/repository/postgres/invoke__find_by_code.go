package postgres

import (
	"context"
	"github.com/doug-martin/goqu"
	"github.com/shellrean-playground/hotel-be-hotel-manager-service/domain"
)

func (m mainRepository) FindByCode(ctx context.Context, code string) (hotel domain.Hotel, err error) {
	dataset := m.db.From(m.tableName).
		Where(goqu.Ex{
			"code":       code,
			"is_deleted": 0,
		})
	_, err = dataset.ScanStructContext(ctx, &hotel)
	return
}
