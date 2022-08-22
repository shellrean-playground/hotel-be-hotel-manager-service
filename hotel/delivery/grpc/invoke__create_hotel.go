package grpc

import (
	"context"
	pb "github.com/shellrean-playground/hotel-be-contract/hotel"
	"github.com/shellrean-playground/hotel-be-hotel-manager-service/dto"
)

func (m mainGrpcDelivery) CreateHotel(ctx context.Context, rpcData *pb.HotelCreate) (*pb.HotelResponse, error) {
	createHotelRequest := dto.CreateHotelRequest{
		Code:    rpcData.Code,
		Name:    rpcData.Name,
		Address: rpcData.Address,
		Rate:    int8(rpcData.Rate),
	}

	svcResult := m.hotelUseCase.CreateHotel(ctx, createHotelRequest)
	return &pb.HotelResponse{
		Code:        svcResult.Code,
		Description: svcResult.Message,
	}, nil
}
