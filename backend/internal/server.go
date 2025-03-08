package main

import (
	"log"
	"net"

	auth "github.com/kiioong/are_they_playing/gen/go/kiioong/authentication"
	authenticationService "github.com/kiioong/are_they_playing/internal/AuthenticationService"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(authenticationService.UnaryInterceptor))

	grpcServer := grpc.NewServer(opts...)
	auth.RegisterAuthenticationServer(grpcServer, authenticationService.NewServer())

	grpcServer.Serve(lis)
}
