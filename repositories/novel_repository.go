package repositories

import "app/models"

type NovelRepository interface {
	FindById(id uint64) *models.Novel
	// FindList(*models.Novel, int limit int offset) map[uint64]*models.Novel
	// Create(*models.Novel)
	// Update(*models.Novel)
	// Delete(*models.Novel)
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

// func (r *novelRepository) FindList(novel *models.Novel, int limit, int offset) map[uint64]*models.Novel {
// 	reuslt := r.DB.Where("").Find(&models.Novel).Limit(limit).Offset(offset)
// 	errors.Is(reuslt.Error, gorm.ErrRecordNotFound)
// }

func (r *novelRepository) Create(novel *models.Novel) (*models.Novel, error) {

	if err := r.DB.Create(&novel).Error; err != nil {
		return nil, err
	}
	return novel, nil
}

func (r *novelRepository) Update(novel *models.Novel) (*models.Novel, error) {
	if err := r.DB.Model(novel).Updates(novel); err != nil {
		return nil, err.Error
	}
	return novel, nil
}

func (r *novelRepository) Delete(id uint64) (ok bool) {

	return ok
}
