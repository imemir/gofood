package order

import (
	"context"
	"encoding/json"
	"github.com/adjust/rmq/v4"
	"github.com/go-redis/redis/v8"
	"github.com/imemir/gofood/pkg/dispose"
	"github.com/imemir/gofood/services/processing/internal/pkg/repository"
	log "github.com/sirupsen/logrus"
	"time"
)

var Consumer consumer

func init() {
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

	queue, err := redisConnection.OpenQueue("order")
	if err != nil {
		log.WithError(err).Fatal("can not open order queue")
	}

	Consumer = consumer{
		queue: queue,
	}
}

type consumer struct {
	queue rmq.Queue
}

func (c consumer) Consume(delivery rmq.Delivery) {
	var order repository.Order
	if err := json.Unmarshal([]byte(delivery.Payload()), &order); err != nil {
		_ = delivery.Reject()
		return
	}

	err := repository.Orders.Save(&order)
	if err != nil {
		_ = delivery.Reject()
		return
	}

	_ = delivery.Ack()
}

func (c consumer) Start() {
	if err := c.queue.StartConsuming(10, time.Second); err != nil {
		log.WithError(err).Fatal("can not start consuming")
	}
	orderConsumer := &consumer{}
	_, err := c.queue.AddConsumer("order-consumer", orderConsumer)
	if err != nil {
		log.WithError(err).Fatal("can not add order consumer")
	}
}
