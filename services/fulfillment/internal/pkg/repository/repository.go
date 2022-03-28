package repository

import (
	"context"
	"github.com/adjust/rmq/v4"
	"github.com/go-redis/redis/v8"
	"github.com/imemir/gofood/pkg/dispose"
	"github.com/imemir/gofood/pkg/envext"
	log "github.com/sirupsen/logrus"
)

type Configs struct {
	RedisAddress  string `env:"REDIS_ADDRESS,required"`
	RedisPassword string `env:"REDIS_PASSWORD,required,file"`
	RedisDB       int    `env:"REDIS_DB,required"`
}

var (
	Orders *orderRepository
)

func init() {
	configs := new(Configs)
	if err := envext.Load(configs); err != nil {
		log.WithError(err).Fatal("can not load repository configs")
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     configs.RedisAddress,
		Password: configs.RedisPassword,
		DB:       configs.RedisDB,
	})
	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		log.WithError(err).Fatal("can not connect to redis")
	}
	dispose.Add(redisClient.Close)
	errChan := make(chan error)
	redisConnection, err := rmq.OpenConnectionWithRedisClient("order", redisClient, errChan)
	if err != nil {
		log.WithError(err).Fatal("can not open rmq connection")
	}

	orderQueue, err := redisConnection.OpenQueue("order")
	if err != nil {
		log.WithError(err).Fatal("can not open order queue")
	}

	Orders = &orderRepository{
		queue: orderQueue,
	}
}
