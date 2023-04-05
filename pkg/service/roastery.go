package service

import (
	"coffeecatalogue4"
	"coffeecatalogue4/pkg/repository"
)

type RoasteryService struct {
	repo repository.Roastery
}

func NewRoasteryService(repo repository.Roastery) *RoasteryService {
	return &RoasteryService{repo: repo}
}

func (s *RoasteryService) Create(userId int, roastery coffeecatalogue4.Roastery) (int, error) {
	return s.repo.Create(userId, roastery)
}
