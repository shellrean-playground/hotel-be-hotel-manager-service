package domain

import (
	"context"
	"github.com/shellrean-playground/hotel-be-hotel-manager-service/dto"
)

type Hotel struct {
	ID        string `db:"id"`
	Code      string `db:"code"`
	Name      string `db:"name"`
	Address   string `db:"address"`
	Rate      int8   `db:"rate"`
	IsDeleted int8   `db:"is_deleted"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

type HotelRepository interface {
	Find(ctx context.Context, hotelId string) (Hotel, error)
	Save(ctx context.Context, hotel *Hotel) error
	Update(ctx context.Context, hotel *Hotel) error
	SoftDelete(ctx context.Context, hotelId string) error
	FindByCode(ctx context.Context, code string) (Hotel, error)
}

type HotelUseCase interface {
	CreateHotel(ctx context.Context, req dto.CreateHotelRequest) dto.SvcResponse
}
