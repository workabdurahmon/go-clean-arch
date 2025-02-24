package service

import (
	"go-clean-arch/internal/repository"
)

type Service struct {
	postgres repository.Repository
	redis    repository.Redis
}

func NewService(repo repository.Repository, redis repository.Redis) (*Service, error) {
	return &Service{
		postgres: repo,
		redis:    redis,
	}, nil
}

func (s *Service) GetArticleByAuthorID() {
	// some logic here with postgres and redis
}
