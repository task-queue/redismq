package redismq

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/redis.v5"
)

type BrokerRedis struct {
	client *redis.Client

	id int64

	consumeQueue   string
	unackedQueue   string
	heartbeatQueue string
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
	c.heartbeatQueue = "redismq::" + queue + "::heartbeat"

	go c.heartbeat()
	go c.observer()

	return nil
}

func (c BrokerRedis) Pop() ([]byte, error) {
	body, err := c.client.BRPopLPush(c.consumeQueue, c.unackedQueue, 0).Result()
	return []byte(body), err
}

func (c BrokerRedis) Ack() {
	c.client.RPop(c.unackedQueue)
}

func (c BrokerRedis) heartbeat() {
	id := strconv.FormatInt(c.id, 10)

	for {
		z := redis.Z{
			Score:  float64(time.Now().Unix()),
			Member: id,
		}

		c.client.ZAdd(c.heartbeatQueue, z)
		time.Sleep(3 * time.Second)
	}
}

func (c BrokerRedis) observer() {

	consumerDieSeconds := int64(15)

	for {
		z := redis.ZRangeBy{
			Max:    strconv.FormatInt(time.Now().Unix()-consumerDieSeconds, 10),
			Offset: 0,
			Count:  -1,
		}
		fmt.Println("Found fails", z)

		res, err := c.client.ZRangeByScore(c.heartbeatQueue, z).Result()
		if err != nil {
			panic(err)
		}

		fmt.Println("Res", res)
		for _, consumerID := range res {
			failQueue := "redismq::" + c.consumeQueue + "::unacked::" + consumerID
			_, err := c.client.RPopLPush(failQueue, c.consumeQueue).Result()
			if err != nil {
				fmt.Printf("\033[0;31m%v\033[0m\n", err)
				// panic(err)
			}
			fmt.Println("REMOVE CONSUMER", consumerID)
			c.client.ZRem(c.heartbeatQueue, consumerID)
		}

		time.Sleep(15 * time.Second)
	}
}
