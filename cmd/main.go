package main

import (
	"go-clean-arch/config"
	"go-clean-arch/internal/repository/postgres"
	"go-clean-arch/internal/repository/redis"
	"go-clean-arch/internal/service"
)

func main() {
	cfg := config.NewConfig()

	// postgres
	repository, err := postgres.NewPostgresRepo(cfg)
	if err != nil {
		panic(err)
	}

	//redis
	redis, err := redis.NewRedisRepo(cfg)
	if err != nil {
		panic(err)
	}

	//service
	service, err := service.NewService(repository, redis)
	if err != nil {
		panic(err)
	}

	service.GetArticleByAuthorID()
}
