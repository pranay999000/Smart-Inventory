package pkg

import (
	"context"
	"net/http"

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

type contextKey string
const ClaimsKey contextKey = "claims"

func WithClaims(r *http.Request, claims *Claims) *http.Request {
	ctx := context.WithValue(r.Context(), ClaimsKey, claims)
	return r.WithContext(ctx)
}

func GetClaims(r *http.Request) (*Claims, bool) {
	claims, ok := r.Context().Value(ClaimsKey).(*Claims)
	return claims, ok
}