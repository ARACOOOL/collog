package http

import "github.com/dgrijalva/jwt-go"

func verifyToken(tokenString string, signingKey []byte) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims, err
}
