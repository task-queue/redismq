package main

import (
	"github.com/task-queue/redismq"
)

func main() {

	r := redismq.New(nil)
	err := r.Connect()

	if err != nil {
		panic(err)
	}
}
