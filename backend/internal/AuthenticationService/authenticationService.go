package authenticationService

import (
	"context"

	as "github.com/kiioong/are_they_playing/gen/go/kiioong/authentication_service"
)

type AuthentificationServer struct {
	as.UnimplementedAuthenticationServer
}

func (s *AuthentificationServer) Login(ctx context.Context, in *as.AuthenticationData) (*as.AuthenticationStatus, error) {
	if in.Username == "Admin" {
		return &as.AuthenticationStatus{
			IsLoggedIn: true,
		}, nil
	}

	return &as.AuthenticationStatus{
		IsLoggedIn: false,
	}, nil
}

func (s *AuthentificationServer) Logout(ctx context.Context, in *as.AuthenticationData) (*as.AuthenticationStatus, error) {
	return &as.AuthenticationStatus{
		IsLoggedIn: false,
	}, nil
}

func NewServer() *AuthentificationServer {
	s := &AuthentificationServer{}
	return s
}
