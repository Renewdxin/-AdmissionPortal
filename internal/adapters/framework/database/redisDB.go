package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() *RedisClient {
	return &RedisClient{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}

func (client *RedisClient) CloseConnection() error {
	if err := client.client.Close(); err != nil {
		log.Fatalln("failed to close the redis service")
		return err
	}
	return nil
}

func (client *RedisClient) SaveVerificationCode(email, code string) error {
	customCtx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	if err := client.client.Set(customCtx, email, code, 0).Err(); err != nil {
		log.Fatalln("failed to save the verification code")
		return err
	}
	return nil
}

func (client *RedisClient) GetVerificationCode(email string) (string, error) {
	customCtx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	code, err := client.client.Get(customCtx, email).Result()
	if err != nil {
		log.Fatalln("failed to get the verification code")
		return "", err
	}
	return code, nil
}
