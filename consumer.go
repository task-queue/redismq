package redismq

import (
	"strconv"

	"gopkg.in/redis.v5"
)

type consumer struct {
	ID int64

	QueueName    string
	UnackedQueeu string

	redis *redis.Client
}

func newConsumer(r *redis.Client, queueName string) *consumer {

	// get unique id for consumer to queue
	id, err := r.Incr("redismq::" + queueName + "::consumers").Result()
	if err != nil {
		panic(err)
	}

	c := &consumer{
		ID: id,

		QueueName:    queueName,
		UnackedQueeu: "redismq::" + queueName + "::unacked::" + strconv.FormatInt(id, 10),

		redis: r,
	}

	return c
}

func (c *consumer) Listen(fn func(b []byte)) {
	for {
		res, err := c.redis.BRPopLPush(c.QueueName, c.UnackedQueeu, 0).Result()

		if err != nil {
			panic(err)
		}

		fn([]byte(res))

		c.redis.RPop(c.UnackedQueeu)
	}
}
