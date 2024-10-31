package vc

import (
	redisClient "github.com/mohamed-gudle/backend-projects/weather-api/internal/redis"
	"github.com/redis/go-redis/v9"
)

type IVCRepository interface {
	GetWeather(location string) (string, error)
	SetWeather(location string, data string) error
}

type RedisVCRepository struct {
	redisClient *redis.Client
}



func NewRedisVCRepository() *RedisVCRepository {
	return &RedisVCRepository{
		redisClient: redisClient.NewRedisClient(),
	}
}

func (r *RedisVCRepository) GetWeather(location string) (string, error) {
	weatherString, err := r.redisClient.Get(ctx, location).Result()

	return weatherString, err
}

func (r *RedisVCRepository) SetWeather(location string, data string) error {
	err := r.redisClient.Set(ctx, location, data, 0).Err()

	return err
}


