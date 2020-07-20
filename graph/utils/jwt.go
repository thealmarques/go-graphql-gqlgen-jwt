package utils

import (
	"go-graphql-jwt/graph/model"

	"github.com/dgrijalva/jwt-go"
)

var issuer = []byte("github/thealmarques")

// DecodeJwt decode jwt
func DecodeJwt(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &model.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return issuer, nil
	})
}

// GenerateJwt create jwt
func GenerateJwt(userID int64, expiredAt int64) string {
	claims := model.UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer:    string(issuer),
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString(issuer)

	return signedToken
}
