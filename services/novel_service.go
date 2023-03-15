package services

import (
	"app/models"
	"app/repositories"
)

type novelSerivce struct {
}

var NovelSerivce = newNovelService()

func newNovelService() *novelSerivce {
	return &novelSerivce{}
}

func (s *novelSerivce) Add(a int) {
	novel := &models.Novel{}
	repositories.NovelRepository.Create(novel)
}
