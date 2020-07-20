package middlewares

import (
	"context"
	"go-graphql-jwt/graph/model"
	"go-graphql-jwt/graph/utils"
	"net"
	"net/http"
	"strings"
)

// UserAuth - user auth middleware structure
type UserAuth struct {
	UserID    int64
	Roles     []string
	IPAddress string
	Token     string
}

var userCtxKey = &contextKey{name: "user"}

type contextKey struct {
	name string
}

// JwtMiddleware middleware for http server
func JwtMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := TokenFromHTTPRequestgo(r)

			userID := UserIDFromHTTPRequestgo(token)
			ip, _, _ := net.SplitHostPort(r.RemoteAddr)
			userAuth := UserAuth{
				UserID:    userID,
				IPAddress: ip,
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &userAuth)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// TokenFromHTTPRequestgo - get jwt token from request
func TokenFromHTTPRequestgo(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	var tokenString string

	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) > 1 {
		tokenString = splitToken[1]
	}
	return tokenString
}

// UserIDFromHTTPRequestgo - get user id from request
func UserIDFromHTTPRequestgo(tokenString string) int64 {
	token, err := utils.DecodeJwt(tokenString)
	if err != nil {
		return 0
	}
	if claims, ok := token.Claims.(*model.UserClaims); ok && token.Valid {
		if claims == nil {
			return 0
		}
		return claims.UserID
	}
	return 0
}

// GetAuthFromContext - Gets context
func GetAuthFromContext(ctx context.Context) *UserAuth {
	raw := ctx.Value(userCtxKey)
	if raw == nil {
		return nil
	}

	return raw.(*UserAuth)
}
