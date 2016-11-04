package redismq

import (
	"fmt"

	"gopkg.in/redis.v5"
)

type RedisMQ struct {
	config *Config
	redis  *redis.Client
}

func New(config *Config) *RedisMQ {
	return &RedisMQ{config: config}
}

func (r *RedisMQ) Connect() error {

	r.redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := r.redis.Ping().Result()

	return err
}

// Declare Queue
func (r *RedisMQ) Declare(queueName string) {

}

func (r *RedisMQ) Publish(queueName string, body []byte) {
	cmd := r.redis.LPush(queueName, string(body))
	_, err := cmd.Result()

	if err != nil {
		panic(err)
	}
}

func (r *RedisMQ) Consume(queueName string, callback func(body []byte)) {
	// register new consumer
	c := newConsumer(r.redis, queueName)
	fmt.Println("Consumer:", c.ID)

	c.Listen(callback)
}
