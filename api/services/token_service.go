package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/technodeguy/real-estate/api/config"
)

type claims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

type TokenService struct {
	jwtConfig   *config.JwtConfig
	redisClient *redis.Client
}

type ITokenService interface {
	Get(hmName string, id string) (string, error)
	Decode(token string) (*claims, error)
	CreateAndSave(userId int, jwtType string) (string, error)
}

func NewTokenService(jwtConfig *config.JwtConfig, redisClient *redis.Client) *TokenService {
	return &TokenService{jwtConfig, redisClient}
}

func (ts *TokenService) save(hmName string, id string, token string) error {
	return ts.redisClient.HMSet(hmName, map[string]interface{}{id: token}).Err()
}

func (ts *TokenService) Get(hmName string, id string) (string, error) {
	data, err := ts.redisClient.HMGet(hmName, id).Result()

	if err != nil || err == redis.Nil {
		return "", err
	}

	if val, ok := data[0].(string); ok {
		return val, nil
	}

	return "", errors.New("")
}

func (ts *TokenService) generate(userId int) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		Id: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ts.jwtConfig.AccessToken.Exp * time.Hour).Unix(),
		},
	}).SignedString([]byte(ts.jwtConfig.AccessToken.Secret))

}

func (ts *TokenService) Decode(token string) (*claims, error) {
	data := &claims{}

	value, err := jwt.ParseWithClaims(token, data, func(token *jwt.Token) (interface{}, error) {
		return []byte(ts.jwtConfig.AccessToken.Secret), nil
	})

	if err != nil {
		return data, err
	}

	if !value.Valid {
		return data, errors.New("")
	}

	return data, err
}

func (ts *TokenService) CreateAndSave(userId int, jwtType string) (string, error) {
	token, err := ts.generate(userId)
	if err != nil {
		return "", err
	}

	err = ts.save(jwtType, strconv.Itoa(userId), token)

	if err != nil {
		return "", err
	}

	return token, nil
}
