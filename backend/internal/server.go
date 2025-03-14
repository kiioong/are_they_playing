package main

import (
	"fmt"
	"log"
	"net"
	"os"

	auth "github.com/kiioong/are_they_playing/gen/go/kiioong/authentication"
	authenticationService "github.com/kiioong/are_they_playing/internal/AuthenticationService"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

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
