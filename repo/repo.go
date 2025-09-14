package repo

import (
	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db             *sqlx.DB
	BrandsRepo     *BrandsRepo
	CategoriesRepo *CategoriesRepo
	ProductRepo    *ProductRepo
	UserRepo       *UserRepo
	FavoriesRepo   *FavoriesRepo
	ReviewsRepo    *ReviewsRepo
}

func NewRepo(db *sqlx.DB) *Repo {

	brandRepo := NewBrandsRepo(db)
	categoriesRepo := NewCategoriesRepo(db)
	productSpecsRepo := NewProductSpecsRepo(db)
	productRepo := NewProductRepo(db, productSpecsRepo, brandRepo, categoriesRepo)
	userRepo := NewUserRepo(db)
	favoriesRepo := NewFavoriesRepo(db)
	reviewsRepo := NewReviewRepo(db)

	return &Repo{
		db: db,

		BrandsRepo:     brandRepo,
		CategoriesRepo: categoriesRepo,
		ProductRepo:    productRepo,
		UserRepo:       userRepo,
		FavoriesRepo:   favoriesRepo,
		ReviewsRepo:    reviewsRepo,
	}
}
