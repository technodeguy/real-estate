package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/technodeguy/real-estate/api/config"
)

type Claims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

type TokenService struct {
	jwtConfig   *config.JwtConfig
	redisClient *redis.Client
}

func (ts *TokenService) Save(hmName string, id string, token string) error {
	return ts.redisClient.HMSet(hmName, map[string]interface{}{id: token}).Err()
}

func (ts *TokenService) Get(hmName string, id string) (interface{}, error) {
	return ts.redisClient.HMGet(hmName, id).Result()
}

func (ts *TokenService) createAndSave(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Id: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	})

	token.SignedString([]byte(ts.jwtConfig.AccessToken.Secret))

	return "", nil
}
