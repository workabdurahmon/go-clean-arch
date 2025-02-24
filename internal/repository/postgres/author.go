package postgres

import (
	"database/sql"
)

type AuthorRepo struct {
	DB *sql.DB
}

func NewAuthorRepo(db *sql.DB) *AuthorRepo {
	return &AuthorRepo{
		DB: db,
	}
}

func (a *AuthorRepo) Create() {
	// create author
}

func (a *AuthorRepo) GetByID() {
	// get author by id
}

func (a *AuthorRepo) GetList() {
	// get list authors
}

func (a *AuthorRepo) Update() {
	// update author
}

func (a *AuthorRepo) Delete() {
	// delete author
}
