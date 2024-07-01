package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var hmacSecret = os.Getenv("SUPABASE_JWT_SECRET")

type UserIdCtxKey string

type Claims struct {
	Email  string `json:"email"`
	UserId string `json:"sub"`
	jwt.RegisteredClaims
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 Unauthorized\n"))
			return
		}

		tokenSegments := strings.Split(token, " ")

		if len(tokenSegments) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 Unauthorized\n"))
			return
		}

		userId, err := parseJWTToken(strings.TrimSpace(tokenSegments[1]), []byte(hmacSecret))

		if err != nil {
			log.Printf("Error parsing token: %s", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 Unauthorized\n"))
			return
		}

		log.Printf("Received request from %s", userId)

		ctx := context.WithValue(r.Context(), UserIdCtxKey("userId"), userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func parseJWTToken(token string, hmacSecret []byte) (userId string, err error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	if err != nil {
		return "", fmt.Errorf("error validating token: %v", err)
	} else if claims, ok := t.Claims.(*Claims); ok {
		fmt.Println(claims)
		return claims.UserId, nil
	}

	return "", fmt.Errorf("error parsing token: %v", err)
}
