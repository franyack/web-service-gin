package repository

import (
	"context"
	"example/web-service-gin/internal/business/domain"
	"example/web-service-gin/internal/business/gateway"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
	"time"
)

func NewRedisAuthRepository() gateway.AuthRepository {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	return &redisAuthRepository{
		client: client,
		ctx:    ctx,
	}
}

type redisAuthRepository struct {
	client *redis.Client
	ctx    context.Context
}

func (repository *redisAuthRepository) CreateAuth(userid uint64, tokenDetails *domain.TokenDetails) error {
	at := time.Unix(tokenDetails.AccessTokenExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(tokenDetails.RefreshTokenExpires, 0)
	now := time.Now()

	errAccess := repository.client.Set(repository.ctx, tokenDetails.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := repository.client.Set(repository.ctx, tokenDetails.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func (repository *redisAuthRepository) FetchAuth(authD *domain.AccessDetails) (uint64, error) {
	userid, err := repository.client.Get(repository.ctx, authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

func (repository *redisAuthRepository) DeleteAuth(givenUuid string) (int64, error) {
	deleted, err := repository.client.Del(repository.ctx, givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
