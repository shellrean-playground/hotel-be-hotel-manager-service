package appioc

import (
	_ "github.com/lib/pq"
	"github.com/shellrean-playground/hotel-be-common-util/provider"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
)

func InitializeServer() {
	appConfiguration := provider.GetConfiguration()
	databaseConnection := provider.GetDatabaseConnection(appConfiguration)

	listener, _ := net.Listen("tcp", appConfiguration.Server.Host+":"+appConfiguration.Server.Port)
	grpcServer := grpc.NewServer()

	loadServiceContainer(appConfiguration, databaseConnection, grpcServer)

	err := grpcServer.Serve(listener)
	if err != nil {
		log.Error("Error when run grpc: ", err.Error())
		os.Exit(1)
	}
}
