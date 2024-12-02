package handlers

import (
	"jwt-auth-service/internal/jwt"
	"jwt-auth-service/internal/service"
)

type Handlers struct {
	s            *service.Service
	tokenManager jwt.TokenManager
}

func NewHandlers(s *service.Service, tokenManager jwt.TokenManager) *Handlers {
	return &Handlers{s: s, tokenManager: tokenManager}
}
