package token

import (
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	*jwt.StandardClaims

	ID    uint   `json:"id"`
	Login string `json:"login"`
	Role  uint8  `json:"role"`
}
