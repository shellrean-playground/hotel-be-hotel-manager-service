package usecase

import (
	"github.com/shellrean-playground/hotel-be-hotel-manager-service/config"
	"github.com/shellrean-playground/hotel-be-hotel-manager-service/domain"
)

type mainUseCase struct {
	hotelRepository domain.HotelRepository
	config          *config.Config
}

func New(args0 domain.HotelRepository, args1 *config.Config) domain.HotelUseCase {
	return &mainUseCase{hotelRepository: args0, config: args1}
}
