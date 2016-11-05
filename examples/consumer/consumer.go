package main

import (
	"flag"
	"fmt"
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

	r := redismq.New(nil)

	err := r.Connect()
	if err != nil {
		panic(err)
	}

	r.Consume(*queueName, func(body []byte) error {
		fmt.Print("receive body", string(body), "...")
		if *sleepMilleseconds > 0 {
			time.Sleep(time.Duration(*sleepMilleseconds) * time.Millisecond)
		}
		fmt.Println("done")
		return nil
	})

}
