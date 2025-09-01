package categories

import (
	"Store-Dio/models"
	"Store-Dio/repo"
)

type CategoriesService struct {
	CategoriesRepo *repo.CategoriesRepo
}

func NewCategoriesService(categoriesRepo *repo.CategoriesRepo) *CategoriesService {
	return &CategoriesService{
		CategoriesRepo: categoriesRepo,
	}
}

func (cs *CategoriesService) InsertCategoriesData(cat models.Category) (bool, error) {
	err := cs.CategoriesRepo.InsertCategoriesRecursive(cat)

	if err != nil {
		return false, err
	}
	return true, nil
}
func (cs *CategoriesService) GetAllCategoriesID() (models.Categories, error) {

	categories, err := cs.CategoriesRepo.GetAllCategoriesID()

	if err != nil {
		return nil, err
	}
	return categories, nil
}
