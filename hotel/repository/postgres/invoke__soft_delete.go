package postgres

import (
	"context"
	"github.com/doug-martin/goqu"
)

func (m mainRepository) SoftDelete(ctx context.Context, hotelId string) (err error) {
	crudExec := m.db.From(m.tableName).
		Where(goqu.Ex{"id": hotelId}).
		Update(goqu.Record{
			"is_deleted": 1,
		})
	_, err = crudExec.ExecContext(ctx)
	return
}
