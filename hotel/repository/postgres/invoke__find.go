package postgres

import (
	"context"
	"github.com/doug-martin/goqu"
	"github.com/shellrean-playground/hotel-be-hotel-manager-service/domain"
)

func (m mainRepository) Find(ctx context.Context, hotelId string) (hotel domain.Hotel, err error) {
	dataset := m.db.From(m.tableName).
		Where(goqu.Ex{
			"id":         hotelId,
			"is_deleted": 0,
		})
	_, err = dataset.ScanStructContext(ctx, &hotel)
	return
}
