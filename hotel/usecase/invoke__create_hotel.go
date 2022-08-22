package usecase

import (
	"context"
	"github.com/shellrean-playground/hotel-be-hotel-manager-service/domain"
	"github.com/shellrean-playground/hotel-be-hotel-manager-service/dto"
	"github.com/shellrean-playground/hotel-be-hotel-manager-service/provider/constant"
	"time"
)

func (m mainUseCase) CreateHotel(ctx context.Context, req dto.CreateHotelRequest) dto.SvcResponse {
	c, cancel := context.WithTimeout(ctx, time.Duration(m.config.ContextTimeout)*time.Second)
	defer cancel()

	hotel, err := m.hotelRepository.FindByCode(c, req.Code)
	if err != nil {
		return dto.SvcResponse{Error: true, Code: constant.ERROR_GENERAL, Message: err.Error()}
	}
	if hotel != (domain.Hotel{}) {
		return dto.SvcResponse{Error: true, Code: constant.DUPLICATION_DATA, Message: "hotel code must be unique"}
	}
	if !(req.Rate <= 5) && !(req.Rate >= 1) {
		return dto.SvcResponse{Error: true, Code: constant.INVALID_PARAM, Message: "rate range 1 - 5"}
	}

	var hotelPersistence domain.Hotel
	hotelPersistence.Code = req.Code
	hotelPersistence.Name = req.Name
	hotelPersistence.Address = req.Address
	hotelPersistence.Rate = req.Rate

	err = m.hotelRepository.Save(c, &hotelPersistence)
	if err != nil {
		return dto.SvcResponse{Error: true, Code: constant.ERROR_GENERAL, Message: err.Error()}
	}
	return dto.SvcResponse{Error: false, Code: constant.SUCCESS, Message: "APPROVED"}
}
