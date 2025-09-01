package brands

import (
	"Store-Dio/models"
	"Store-Dio/repo"
	"fmt"
)

type BrandsService struct {
	BrandsRepo *repo.BrandsRepo
}

func NewBrandsService(brandsRepo *repo.BrandsRepo) *BrandsService {
	return &BrandsService{
		BrandsRepo: brandsRepo,
	}
}

func (bs *BrandsService) InsertBrandData(brands *models.Brands) error {
	if brands.Brand == nil {
		return fmt.Errorf("Invalid data")
	}
	err := bs.BrandsRepo.InsertBrandData(brands)

	if err != nil {
		return fmt.Errorf("Error : %w", err.Error())
	}
	return nil
}
