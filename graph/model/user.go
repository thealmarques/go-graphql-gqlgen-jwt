package model

import (
	"github.com/dgrijalva/jwt-go"
)

// User model
type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

// UserClaims for jwt
type UserClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}
