package main

import (
	"flag"
	"strconv"
	"time"

	"github.com/task-queue/redismq"
)

var queueName *string
var sleepMilleseconds *int

func init() {
	queueName = flag.String("queue", "task", "a queue name")
	sleepMilleseconds = flag.Int("sleep", 100, "Sleep in millesonds between messages")
}

func main() {

	flag.Parse()

	config := &redismq.Config{
		RedisDSN: "localhost:6379",
	}

	r := redismq.New(config)

	err := r.Connect()
	if err != nil {
		panic(err)
	}

	i := 0

	for true {
		i++
		r.Publish(*queueName, []byte("Ping command "+strconv.Itoa(i)))
		if *sleepMilleseconds > 0 {
			time.Sleep(time.Duration(*sleepMilleseconds) * time.Millisecond)
		}
	}
}
