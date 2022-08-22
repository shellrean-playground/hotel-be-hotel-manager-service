package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	pb "github.com/shellrean-playground/hotel-be-contract/hotel"
	"github.com/shellrean-playground/hotel-be-hotel-manager-service/config"
	_hotelRpcService "github.com/shellrean-playground/hotel-be-hotel-manager-service/hotel/delivery/grpc"
	_hotelRepository "github.com/shellrean-playground/hotel-be-hotel-manager-service/hotel/repository/postgres"
	_hotelUseCase "github.com/shellrean-playground/hotel-be-hotel-manager-service/hotel/usecase"
	"google.golang.org/grpc"
	"net"
)

func main() {
	lis, _ := net.Listen("tcp", ":1200")
	s := grpc.NewServer()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		"localhost",
		"postgres",
		"postgres",
		"exo_hotel",
		"5432",
		"Asia/Jakarta",
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	hotelRepository := _hotelRepository.New(db)
	hotelUseCase := _hotelUseCase.New(hotelRepository, &config.Config{ContextTimeout: 30})
	hotelRpcService := _hotelRpcService.New(hotelUseCase)

	pb.RegisterHotelServer(s, hotelRpcService)
	s.Serve(lis)
}
