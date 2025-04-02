package main

import (
	"fmt"
	"log"
	"net"
	"os"

	auth "github.com/kiioong/are_they_playing/gen/go/kiioong/authentication"
	league_management_proto "github.com/kiioong/are_they_playing/gen/go/kiioong/league_management"
	authenticationService "github.com/kiioong/are_they_playing/internal/AuthenticationService"
	"github.com/kiioong/are_they_playing/internal/Database"
	leaguemanagement "github.com/kiioong/are_they_playing/internal/LeagueManagemant"
	"google.golang.org/grpc"
)

func main() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	Database.InitDatabase(dsn)

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(authenticationService.UnaryInterceptor), grpc.StreamInterceptor(authenticationService.StreamInterceptor))

	grpcServer := grpc.NewServer(opts...)
	auth.RegisterAuthenticationServer(grpcServer, authenticationService.NewServer())
	league_management_proto.RegisterLeagueManagementServer(grpcServer, leaguemanagement.NewServer())

	grpcServer.Serve(lis)
}
