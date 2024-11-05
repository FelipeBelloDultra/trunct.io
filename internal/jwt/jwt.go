package jwt

import (
	"fmt"
	"os"
	"time"

	_jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	jwtSecretKey   = []byte(os.Getenv("JWT_SECRET_KEY"))
	expirationTime = time.Now().Add(time.Hour * 24).Unix()
)

func CreateTokenFromID(id uuid.UUID) (string, error) {
	token := _jwt.NewWithClaims(
		_jwt.SigningMethodHS256,
		_jwt.MapClaims{
			"id":  id.String(),
			"exp": expirationTime,
		},
	)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := _jwt.Parse(tokenString, func(token *_jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
