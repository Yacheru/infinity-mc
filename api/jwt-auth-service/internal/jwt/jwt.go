package jwt

import (
	"errors"
	"fmt"
	"jwt-auth-service/internal/entities"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenManager interface {
	NewAccessToken(ipAddr, userId, role string, ttl time.Duration) (string, error)
	NewRefreshToken(ipAddr, userId, role string, ttl time.Duration) (string, error)
	ValidToken(t string) (*entities.Claims, error)
}

type Manager struct {
	salt string
}

func NewJWTManager(salt string) *Manager {
	return &Manager{salt: salt}
}

var (
	TokenMalformed = errors.New("token is malformed")
	TokenExpired   = errors.New("token is expired")
	TokenInvalid   = errors.New("token is invalid")
)

func (m *Manager) NewAccessToken(ipAddr, userId, role string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, entities.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			Subject:   userId,
		},
		IpAddr: ipAddr,
		Role:   role,
	})

	return token.SignedString([]byte(m.salt))
}

func (m *Manager) NewRefreshToken(ipAddr, userId, role string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, entities.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			Subject:   userId,
		},
		IpAddr: ipAddr,
		Role:   role,
	})

	return token.SignedString([]byte(m.salt))
}

func (m *Manager) ValidToken(t string) (*entities.Claims, error) {
	claims := new(entities.Claims)

	_, err := jwt.ParseWithClaims(t, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}
		return []byte(m.salt), nil
	})

	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, TokenExpired
			} else {
				return nil, TokenInvalid
			}
		}
	}

	return claims, nil
}
