package service

import (
	"github.com/yacheru/infinity-mc.ru/backend/internal/app/repository"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response/mc"
)

type McService struct {
	repo repository.McBans
}

func NewMcBansService(repo repository.McBans) *McService {
	return &McService{repo: repo}
}

func (s *McService) GetPunishments(limit int, pType string) ([]mc.LbPunishments, error) {
	return s.repo.GetPunishments(limit, pType)
}
