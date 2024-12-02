package entities

import "github.com/golang-jwt/jwt"

type Claims struct {
	jwt.StandardClaims
	IpAddr string `json:"ip_addr"`
	Role   string `json:"role"`
}
