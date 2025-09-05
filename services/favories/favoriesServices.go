package favories

import (
	"Store-Dio/models"
	"Store-Dio/repo"
	"fmt"
)

type FavoriesService struct {
	FavoriesRepo *repo.FavoriesRepo
}

func NewFavoriesService(repo *repo.FavoriesRepo) *FavoriesService {
	return &FavoriesService{
		FavoriesRepo: repo,
	}
}

func (fs *FavoriesService) AddFavori(data *models.Product, user_id int) error {
	if data.ID == 0 {
		return fmt.Errorf("Invalid data")
	}
	err := fs.FavoriesRepo.AddFavori(data, user_id)
	if err != nil {
		return err
	}
	return nil
}
func (fs *FavoriesService) RemoveFavori(id int, user_id int) error {
	if id == 0 || user_id == 0 {
		return fmt.Errorf("Invalid data")
	}
	var data models.Favori
	data.ID = id
	data.UserID = user_id
	err := fs.FavoriesRepo.RemoveFavori(&data)
	if err != nil {
		return err
	}
	return nil
}
func (fs *FavoriesService) GetFavourites(page int, user_id int) ([]*models.Favori, error) {
	if user_id == 0 {
		return nil, fmt.Errorf("Invalid data")
	}
	favourites, err := fs.FavoriesRepo.GetFavourites(page, user_id)

	if err != nil {
		return nil, err
	}
	return favourites, nil
}
