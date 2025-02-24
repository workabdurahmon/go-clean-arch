package postgres

import (
	"database/sql"
	"go-clean-arch/config"
	"go-clean-arch/internal/repository"

	"log"
)

type PostgresRepo struct {
	DB      *sql.DB
	author  *AuthorRepo
	article *ArticleRepo
}

func NewPostgresRepo(cfg config.Config) (*PostgresRepo, error) {
	db, err := sql.Open(cfg.DB.Driver, cfg.DB.Username+":"+cfg.DB.Password+"@tcp("+cfg.DB.Host+":"+cfg.DB.Port+")/"+cfg.DB.Name)
	if err != nil {
		log.Fatal("failed to open connection to database", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("failed to ping database", err)
		return nil, err
	}

	defer db.Close()

	return &PostgresRepo{
		DB:      db,
		author:  NewAuthorRepo(db),
		article: NewArticleRepo(db),
	}, nil
}

func (p *PostgresRepo) Close() {
	if p.DB != nil {
		p.DB.Close()
	}
}

func (p *PostgresRepo) Author() repository.AuthorRepository {
	return p.author
}

func (p *PostgresRepo) Article() repository.ArticleRepository {
	return p.article
}
