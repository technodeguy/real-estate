package services

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/technodeguy/real-estate/api/config"
)

type Claims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

type tokenService struct {
	jwtConfig   *config.JwtConfig
	redisClient *redis.Client
}

type ITokenService interface {
	Get(hmName string, id string) (interface{}, error)
	CreateAndSave(userId int, jwtType string) (string, error)
}

func NewTokenService(jwtConfig *config.JwtConfig, redisClient *redis.Client) *tokenService {
	return &tokenService{jwtConfig, redisClient}
}

func (ts *tokenService) save(hmName string, id string, token string) error {
	return ts.redisClient.HMSet(hmName, map[string]interface{}{id: token}).Err()
}

func (ts *tokenService) Get(hmName string, id string) (interface{}, error) {
	return ts.redisClient.HMGet(hmName, id).Result()
}

func (ts *tokenService) CreateAndSave(userId int, jwtType string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Id: userId,
		StandardClaims: jwt.StandardClaims{
			// TODO sync with config
			ExpiresAt: time.Now().Add(ts.jwtConfig.AccessToken.Exp * time.Hour).Unix(),
		},
	}).SignedString([]byte(ts.jwtConfig.AccessToken.Secret))

	if err != nil {
		return "", err
	}

	err = ts.save(jwtType, strconv.Itoa(userId), token)

	if err != nil {
		return "", err
	}

	return token, nil
}
