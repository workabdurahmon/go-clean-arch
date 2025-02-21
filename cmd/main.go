package main

import (
	"go-clean-arch/config"
	"go-clean-arch/internal/repository/postgres"
	"go-clean-arch/internal/repository/redis"
	"log"
)

func main() {
	cfg := config.NewConfig()

	// postgres
	postgres, err := postgres.NewPostgres(cfg)
	if err != nil {
		log.Fatal("failed to connect to postgres", err)
	}
	defer postgres.Close()

	//redis
	redis, err := redis.NewRedisRepository(cfg)
	if err != nil {
		log.Fatal("failed to connect to redis", err)
	}
	defer redis.Close()

}
