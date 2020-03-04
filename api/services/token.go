package services

import (
	"github.com/go-redis/redis"
)

type TokenService struct {
	redisClient *redis.Client
}

func (ts *TokenService) Save(hmName string, userId string, token string) error {
	return ts.redisClient.HMSet(hmName, map[string]interface{}{userId: token}).Err()
}

func (ts *TokenService) Get(hmName string, userId string) (interface{}, error) {
	return ts.redisClient.HMGet(hmName, userId).Result()
}
