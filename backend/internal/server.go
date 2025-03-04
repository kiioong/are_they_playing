package main

import (
	"log"
	"net"

	as "github.com/kiioong/are_they_playing/gen/go/kiioong/authentication_service"
	authenticationService "github.com/kiioong/are_they_playing/internal/AuthenticationService"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	as.RegisterAuthentificationServer(grpcServer, authenticationService.NewServer())

	grpcServer.Serve(lis)
}
