package authenticationService

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	auth "github.com/kiioong/are_they_playing/gen/go/kiioong/authentication"
	"github.com/kiioong/are_they_playing/internal/Database"
	hash "github.com/kiioong/are_they_playing/internal/Hash"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Custom error types for better error handling
var (
	ErrMissingMetadata   = status.Error(codes.InvalidArgument, "missing metadata")
	ErrMissingToken      = status.Error(codes.Unauthenticated, "missing authorization token")
	ErrInvalidToken      = status.Error(codes.Unauthenticated, "invalid token format")
	ErrTokenExpired      = status.Error(codes.Unauthenticated, "token has expired")
	ErrInvalidSignature  = status.Error(codes.Unauthenticated, "invalid token signature")
	ErrUserNotFound      = status.Error(codes.NotFound, "user not found")
	ErrInvalidPassword   = status.Error(codes.Unauthenticated, "invalid password")
	ErrInvalidServiceKey = status.Error(codes.PermissionDenied, "invalid service authentication key")
)

var excludedMethods = map[string]bool{
	"/services.Authentication/Login":                       true,
	"/services.Authentication/AuthenticateInternalService": true,
}

type AuthentificationServer struct {
	auth.UnimplementedAuthenticationServer
}

type User struct{}

func (s *AuthentificationServer) Login(ctx context.Context, in *auth.LoginData) (*auth.Session, error) {
	var user Database.User

	result := Database.DB.Where("username = ?", strings.ToLower(in.Username)).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, Database.DB.Error) {
			return nil, ErrUserNotFound
		}
		return nil, status.Error(codes.Internal, fmt.Sprintf("database error: %v", result.Error))
	}

	if !hash.VerifyPassword(in.Password, user.Password) {
		return nil, ErrInvalidPassword
	}

	return &auth.Session{
		JwtToken: GenerateJWT(user.ID),
	}, nil
}

func (s *AuthentificationServer) Logout(ctx context.Context, in *auth.Session) (*auth.Session, error) {
	return &auth.Session{
		JwtToken: "",
	}, nil
}

func (s *AuthentificationServer) AuthenticateInternalService(ctx context.Context, in *auth.ServiceAuthToken) (*auth.Session, error) {
	if in.Token != os.Getenv("INTERNAL_SERVICE_AUTH_KEY") {
		return nil, ErrInvalidServiceKey
	}

	return &auth.Session{
		JwtToken: GenerateJWT(in.ServiceId),
	}, nil
}

func (s *AuthentificationServer) ValidateToken(ctx context.Context, in *auth.Session) (*auth.Session, error) {
	claims, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}

	if claims == nil {
		return nil, ErrInvalidToken
	}

	return &auth.Session{
		JwtToken: in.JwtToken,
	}, nil
}

func NewServer() *AuthentificationServer {
	s := &AuthentificationServer{}
	return s
}

// Secret key for JWT validation (should be stored securely)
var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

// AuthInterceptor checks JWT token in metadata
func AuthInterceptor(ctx context.Context) (context.Context, error) {
	claims, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}

	// Store claims in context for later use
	ctx = context.WithValue(ctx, User{}, claims.Subject)
	return ctx, nil
}

// UnaryInterceptor is a gRPC middleware for authentication
func UnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// Check if the current method is in the excluded list
	if _, excluded := excludedMethods[info.FullMethod]; excluded {
		return handler(ctx, req)
	}

	newCtx, err := AuthInterceptor(ctx)
	if err != nil {
		return nil, err
	}
	return handler(newCtx, req)
}

type wrappedStream interface {
	grpc.ServerStream
	SetContext(context.Context)
}

type wrapper struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrapper) Context() context.Context {
	return w.ctx
}

func (w *wrapper) SetContext(ctx context.Context) {
	w.ctx = ctx
}

func newWrappedStream(s grpc.ServerStream) wrappedStream {
	ctx := s.Context()
	return &wrapper{s, ctx}
}

func StreamInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if _, excluded := excludedMethods[info.FullMethod]; excluded {
		return nil
	}

	wrappedStream := newWrappedStream(
		ss,
	)

	newCtx, err := AuthInterceptor(ss.Context())

	wrappedStream.SetContext(newCtx)

	if err != nil {
		return err
	}

	// authentication (token verification)
	return handler(srv, wrappedStream)
}

func validateToken(ctx context.Context) (*jwt.RegisteredClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrMissingMetadata
	}

	authHeader, exists := md["authorization"]
	if !exists || len(authHeader) == 0 {
		return nil, ErrMissingToken
	}

	tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")
	if tokenString == authHeader[0] { // No "Bearer " prefix
		return nil, ErrInvalidToken
	}

	// Validate token
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, ErrInvalidSignature
		}
		return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("token validation failed: %v", err))
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

func GenerateJWT(user_id uint64) string {
	claims := jwt.RegisteredClaims{
		Subject:   strconv.FormatUint(user_id, 10),                    // User ID
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 24-hour expiry
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return signedToken
}
