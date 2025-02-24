package repository

import "context"

type AuthorRepository interface {
	Create()
	GetByID()
	GetList()
	Update()
	Delete()
}

type ArticleRepository interface {
	Create()
	GetByID()
	GetList()
	Update()
	Delete()
}

type Repository interface {
	Author() AuthorRepository
	Article() ArticleRepository
}

type Redis interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
}
