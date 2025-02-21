package postgres

import (
	"database/sql"
	"go-clean-arch/config"
	"log"
)

type PostgresRepository struct {
	DB      *sql.DB
	Author  *AuthorRepository
	Article *ArticleRepository
}

func NewPostgres(cfg config.Config) (*PostgresRepository, error) {
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

	return &PostgresRepository{
		DB:      db,
		Author:  NewAuthorRepository(db),
		Article: NewArticleRepository(db),
	}, nil
}

func (p *PostgresRepository) Close() {
	if p.DB != nil {
		p.DB.Close()
	}
}
