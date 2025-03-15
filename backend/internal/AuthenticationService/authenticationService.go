package authenticationService

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	auth "github.com/kiioong/are_they_playing/gen/go/kiioong/authentication"
	"github.com/kiioong/are_they_playing/internal/Database"
	hash "github.com/kiioong/are_they_playing/internal/Hash"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Methods that should be excluded from authentication (e.g., Login)
var excludedMethods = map[string]bool{
	"/services.Authentication/Login": true, // Exclude Login RPC
}

type AuthentificationServer struct {
	auth.UnimplementedAuthenticationServer
}

func (s *AuthentificationServer) Login(ctx context.Context, in *auth.LoginData) (*auth.Session, error) {
	var user Database.User

	result := Database.DB.Where("username = ?", strings.ToLower(in.Username)).First(&user)

	if result.Error != nil {
		return &auth.Session{
			JwtToken: "",
		}, nil
	}

	if hash.VerifyPassword(in.Password, user.Password) {
		return &auth.Session{
			JwtToken: GenerateJWT(),
		}, nil
	}

	return &auth.Session{
		JwtToken: "",
	}, nil
}

func (s *AuthentificationServer) Logout(ctx context.Context, in *auth.Session) (*auth.Session, error) {
	return &auth.Session{
		JwtToken: "",
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
	// Extract metadata from context
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("missing metadata")
	}

	// Get authorization header
	authHeader, exists := md["authorization"]
	if !exists || len(authHeader) == 0 {
		return nil, errors.New("missing authorization token")
	}

	// Extract token (format: "Bearer <token>")
	tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")
	if tokenString == authHeader[0] { // No "Bearer " prefix
		return nil, errors.New("invalid token format")
	}

	// Validate token
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	userId := 1

	// Store claims in context for later use
	ctx = context.WithValue(ctx, userId, claims.Subject)
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
		return handler(ctx, req) // Skip auth, proceed with request
	}

	newCtx, err := AuthInterceptor(ctx)
	if err != nil {
		return nil, err
	}
	return handler(newCtx, req)
}

func GenerateJWT() string {
	claims := jwt.RegisteredClaims{
		Subject:   "12345",                                       // User ID
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), // 1-hour expiry
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return signedToken
}
