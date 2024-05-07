package service

import "github.com/yacheru/infinity-mc.ru/backend/pkg/repository"

type McService interface {
}

type Service struct {
	McService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
