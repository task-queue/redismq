package main

import (
	"fmt"

	"github.com/task-queue/redismq"
)

func main() {

	r := redismq.New(nil)
	fmt.Println(r)

}
