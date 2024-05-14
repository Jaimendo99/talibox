package Models

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	Sub   uint `json:"sub"`
	Admin bool `json:"admin"`
	jwt.RegisteredClaims
}
