package jwt

import (
	"errors"
	"os"
	"time"

	_jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	jwtSecretKey    = []byte(os.Getenv("JWT_SECRET_KEY"))
	expirationTime  = time.Now().Add(time.Hour * 24).Unix()
	ErrInvalidToken = errors.New("invalid token")
)

func CreateTokenFromID(id uuid.UUID) (string, error) {
	token := _jwt.NewWithClaims(
		_jwt.SigningMethodHS256,
		_jwt.MapClaims{
			"sub": id.String(),
			"exp": expirationTime,
		},
	)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*_jwt.Token, error) {
	token, err := _jwt.Parse(tokenString, func(token *_jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	return token, nil
}
