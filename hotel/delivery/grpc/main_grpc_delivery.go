package grpc

import (
	pb "github.com/shellrean-playground/hotel-be-contract/hotel"
	"github.com/shellrean-playground/hotel-be-hotel-manager-service/domain"
)

type mainGrpcDelivery struct {
	pb.UnimplementedHotelServer
	hotelUseCase domain.HotelUseCase
}

func New(args0 domain.HotelUseCase) pb.HotelServer {
	return &mainGrpcDelivery{hotelUseCase: args0}
}
