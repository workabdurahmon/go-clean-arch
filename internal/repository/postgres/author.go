package postgres

import (
	"database/sql"
)

type AuthorRepository struct {
	DB *sql.DB
}

func NewAuthorRepository(db *sql.DB) *AuthorRepository {
	return &AuthorRepository{
		DB: db,
	}
}

func (a *AuthorRepository) Create() {
	// create author
}

func (a *AuthorRepository) GetByID() {
	// get author by id
}

func (a *AuthorRepository) GetList() {
	// get list authors
}

func (a *AuthorRepository) Update() {
	// update author
}

func (a *AuthorRepository) Delete() {
	// delete author
}
