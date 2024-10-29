package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pranay999000/smart-inventory/api-gateway/pkg"
)

func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		if authHeader == "" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(struct{
				Success		bool		`json:"success"`
				ErrMessage	string		`json:"err_message"`
			} {
				Success: false,
				ErrMessage: "Authorization header missing",
			})
			return
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]

		claims := &pkg.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			log.Printf("invalid token: %v\n", err)
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(struct{
				Success		bool		`json:"success"`
				ErrMessage	string		`json:"err_message"`
			} {
				Success: false,
				ErrMessage: "Invalid or Expired token",
			})
			return
		}

		r = pkg.WithClaims(r, claims)

		next.ServeHTTP(w, r)
	})

}