package main

import "fmt"

func main() {
	ch := make(chan int)
	go publish(ch) // run in background
	reader(ch)     // run in the main thread

}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("received %d\n", x)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // So that the function reader is aware that the thread won't keep sending data, close the channel. It's prevent the dead lock
}
