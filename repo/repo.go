package repo

import (
	"database/sql"

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
	productRepo := NewProductRepo(db)
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
func (r *Repo) SafeQueryRow(query string, args ...any) *sqlx.Row {
	return r.db.QueryRowx(query, args...)
}
func (r *Repo) SafeQuery(query string, args ...any) (*sqlx.Rows, error) {
	return r.db.Queryx(query, args...)
}
func (r *Repo) SafeExec(query string, args ...any) (sql.Result, error) {
	return r.db.Exec(query, args...)
}
