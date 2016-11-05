package main

import (
	"flag"
	"fmt"

	"github.com/task-queue/redismq"
)

var queueName *string

func init() {
	queueName = flag.String("queue", "task", "a queue name")
}

func main() {

	flag.Parse()

	r := redismq.New(nil)

	err := r.Connect()
	if err != nil {
		panic(err)
	}

	r.Consume(*queueName, func(body []byte) error {
		fmt.Println("receive body", string(body))
		return nil
	})

}
