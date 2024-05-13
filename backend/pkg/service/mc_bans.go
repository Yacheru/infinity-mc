package service

import (
	"github.com/yacheru/infinity-mc.ru/backend"
	"github.com/yacheru/infinity-mc.ru/backend/pkg/repository"
)

type McService struct {
	repo repository.McBans
}

func NewMcBansService(repo repository.McBans) *McService {
	return &McService{repo: repo}
}

func (s *McService) GetAllBans(limit int) ([]backend.LbPunishments, error) {
	return s.repo.GetAllBans(limit)
}
