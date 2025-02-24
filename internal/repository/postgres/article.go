package postgres

import (
	"database/sql"
)

type ArticleRepo struct {
	DB *sql.DB
}

func NewArticleRepo(db *sql.DB) *ArticleRepo {
	return &ArticleRepo{
		DB: db,
	}
}

func (a *ArticleRepo) Create() {
	// create article
}

func (a *ArticleRepo) GetByID() {
	// get article by id
}

func (a *ArticleRepo) GetList() {
	// get list articles
}

func (a *ArticleRepo) Update() {
	// update article
}

func (a *ArticleRepo) Delete() {
	// delete article
}
