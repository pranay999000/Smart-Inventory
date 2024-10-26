package pkg

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserId		string		`json:"user_id"`
	FirstName	string		`json:"first_name"`
	MiddleName	string		`json:"middle_name"`
	LastName	string		`json:"last_name"`
	Avatar		string		`json:"avatar"`
	jwt.RegisteredClaims
}

func GenerateJWTToken(userId string, firstName string, middleName string, lastName string, avatar string) (string, error) {
	claims := &Claims{
		UserId: userId,
		FirstName: firstName,
		MiddleName: middleName,
		LastName: lastName,
		Avatar: avatar,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}