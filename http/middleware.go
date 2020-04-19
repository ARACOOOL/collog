package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func authenticationMiddleware(signingKey string, next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(ServerError{
				Status:  "error",
				Message: "Missing Authorization Header",
			})
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := verifyToken(tokenString, []byte(signingKey))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(ServerError{
				Status:  "error",
				Message: "Error verifying a token: " + err.Error(),
			})
			return
		}

		email := claims.(jwt.MapClaims)["email"]
		if email == nil {
			_ = json.NewEncoder(w).Encode(ServerError{
				Status:  "error",
				Message: "Token is invalid",
			})
			return
		}

		r.Header.Set("email", email.(string))

		next(w, r)
	}
}
