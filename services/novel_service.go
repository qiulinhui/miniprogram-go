package services

import (
	"app/models"
	"app/repositories"
)

type NovelService interface {
	Get(uint64) *models.Novel
	Add(novel *models.Novel) bool
}

type novelSerivce struct {
	novelRepo repositories.NovelRepository
}

func NewNovelService(novelRepo repositories.NovelRepository) *novelSerivce {
	return &novelSerivce{
		novelRepo,
	}
}

func (s *novelSerivce) Get(id uint64) *models.Novel {
	s.novelRepo.FindById(id)
	return s.novelRepo.FindById(id)
}

func (s *novelSerivce) Add(novel *models.Novel) (ok bool) {
	if err := s.novelRepo.Insert(novel); err == nil {
		ok = true
	}
	return ok
}
