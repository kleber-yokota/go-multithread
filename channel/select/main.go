package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id      int64
	message string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	var i int64 = 0

	go func() {
		for {
			atomic.AddInt64(&i, 1) //prevent race condition
			msg := Message{i, "RabbitMQ message"}
			c1 <- msg
		}
	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			msg := Message{i, "Kafka message"}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Printf("message: id: %d - %s,\n", msg.id, msg.message)
		case msg := <-c2:
			fmt.Printf("message: id: %d - %s,\n", msg.id, msg.message)
		case <-time.After(time.Millisecond * 500): // after 500 ms the result of process is timeout
			fmt.Print("timeout\n")
		}
	}

}
