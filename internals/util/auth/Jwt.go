package auth

import (
	"errors"
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtClaims struct {
	Email  string `json:"email"`
	UserId string `json:"userId"`

	jwt.RegisteredClaims
}

type AuthToken struct {
	logger *slog.Logger
	secret []byte
}

func GetNewAuthToken(logger *slog.Logger) *AuthToken {
	// get secret from .env
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET not set")
	}

	return &AuthToken{
		logger: logger,
		secret: []byte(secret),
	}
}

// Generate jwt token
func (at *AuthToken) GenerateJwtToken(email string, userId string) (string, string, error) {
	claims := JwtClaims{
		Email:  email,
		UserId: userId,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(at.secret)
	if err != nil {
		at.logger.Error("failed to sign jwt token", "error", err)
		return "", "", err
	}

	sessionId := uuid.New().String()
	if sessionId == "" {
		at.logger.Error("faild to generate the session Id")
		return "", "", errors.New("faild to generate the session Id")
	}

	return signedToken, sessionId, nil
}

// verify token
func (at *AuthToken) VerifyJwtToken(token string) {

}
