package authenticationService

import (
	"context"
	"fmt"

	as "github.com/kiioong/are_they_playing/gen/go/kiioong/authentication_service"
)

type AuthentificationServer struct {
	as.UnimplementedAuthentificationServer
}

func (s *AuthentificationServer) Login(ctx context.Context, in *as.AuthenticationData) (*as.AuthentificationStatus, error) {
	fmt.Println("Test")
	if in.Username == "Admin" {
		return &as.AuthentificationStatus{
			IsLoggedIn: true,
		}, nil
	}

	return &as.AuthentificationStatus{
		IsLoggedIn: false,
	}, nil
}

func (s *AuthentificationServer) Logout(ctx context.Context, in *as.AuthenticationData) (*as.AuthentificationStatus, error) {
	return &as.AuthentificationStatus{
		IsLoggedIn: false,
	}, nil
}

func NewServer() *AuthentificationServer {
	s := &AuthentificationServer{}
	return s
}
