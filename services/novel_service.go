package services

import (
	"app/models"
	"app/repositories"
)

type NovelService interface {
	Get(uint64) *models.Novel
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
