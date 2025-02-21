package postgres

import (
	"database/sql"
)

type ArticleRepository struct {
	DB *sql.DB
}

func NewArticleRepository(db *sql.DB) *ArticleRepository {
	return &ArticleRepository{
		DB: db,
	}
}

func (a *ArticleRepository) Create() {
	// create article
}

func (a *ArticleRepository) GetByID() {
	// get article by id
}

func (a *ArticleRepository) GetList() {
	// get list articles
}

func (a *ArticleRepository) Update() {
	// update article
}

func (a *ArticleRepository) Delete() {
	// delete article
}
