package redismq

import (
	"strconv"

	"gopkg.in/redis.v5"
)

type BrokerRedis struct {
	client *redis.Client

	id int64

	consumeQueue string
	unackedQueue string
}

func (c BrokerRedis) Push(queue string, body []byte) error {
	cmd := c.client.LPush(queue, string(body))
	_, err := cmd.Result()
	return err
}

func (c *BrokerRedis) InitConsumer(queue string) error {
	var err error

	c.id, err = c.client.Incr("redismq::" + queue + "::consumers").Result()
	if err != nil {
		return err
	}

	c.consumeQueue = queue
	c.unackedQueue = "redismq::" + queue + "::unacked::" + strconv.FormatInt(c.id, 10)

	return nil
}

func (c BrokerRedis) Pop() ([]byte, error) {
	body, err := c.client.BRPopLPush(c.consumeQueue, c.unackedQueue, 0).Result()
	return []byte(body), err
}

func (c BrokerRedis) Ack() {
	c.client.RPop(c.unackedQueue)
}
