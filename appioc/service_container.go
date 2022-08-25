package appioc

import (
	"database/sql"
	"github.com/shellrean-playground/hotel-be-common-util/config"
	pb "github.com/shellrean-playground/hotel-be-contract/hotel"
	_hotelRpcService "github.com/shellrean-playground/hotel-be-hotel-manager-service/hotel/delivery/grpc"
	_hotelRepository "github.com/shellrean-playground/hotel-be-hotel-manager-service/hotel/repository/postgres"
	_hotelUseCase "github.com/shellrean-playground/hotel-be-hotel-manager-service/hotel/usecase"
	"google.golang.org/grpc"
)

func loadServiceContainer(appConfig *config.Config, databaseConnection *sql.DB, grpcServer *grpc.Server) {
	hotelRepository := _hotelRepository.New(databaseConnection)
	hotelUseCase := _hotelUseCase.New(hotelRepository, appConfig)
	hotelRpcService := _hotelRpcService.New(hotelUseCase)

	pb.RegisterHotelServer(grpcServer, hotelRpcService)
}
