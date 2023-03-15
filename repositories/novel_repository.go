package repositories

import (
	"app/app"
	"app/models"
)

var NovelRepository = newNovelRepository()

type novelRepository struct {
}

func newNovelRepository() *novelRepository {
	return &novelRepository{}
}

func (r *novelRepository) Find(id int64) *models.Novel {
	ret := &models.Novel{}
	if err := app.DB.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}

func (r *novelRepository) Create(novel *models.Novel) (bool, *models.Novel) {

	if err := app.DB.Create(&novel).Error; err != nil {
		return false, novel
	}
	return true, nil
}
