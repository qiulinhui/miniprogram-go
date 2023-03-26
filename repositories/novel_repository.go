package repositories

import "app/models"

type NovelRepository interface {
	FindById(id uint64) *models.Novel
	// FindList(*models.Novel, int limit int offset) map[uint64]*models.Novel
	Insert(*models.Novel) error
	Update(*models.Novel) error
	Delete(*models.Novel) error
}

type novelRepository struct {
	repository
}

func NewNovelRepository() NovelRepository {
	return &novelRepository{
		*NewRepository(),
	}
}

func (r *novelRepository) FindById(id uint64) *models.Novel {
	model := &models.Novel{}
	if err := r.DB.First(model, "id = ?", id).Error; err != nil {
		return nil
	}
	return model
}

func (r *novelRepository) Insert(novel *models.Novel) error {

	if err := r.DB.Create(&novel).Error; err != nil {
		return err
	}
	return nil
}

func (r *novelRepository) Update(novel *models.Novel) error {
	return r.DB.Model(novel).Updates(novel).Error
}

func (r *novelRepository) Delete(novel *models.Novel) error {
	return r.DB.Delete(&novel).Error
}
