package main

import "fmt"

func worker(workerID int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
	}
}

func main() {
	data := make(chan int)
	for i := 0; i < 1000; i++ {
		go worker(i, data)
	}
	for i := 0; i < 100000; i++ {
		data <- i
	}
}
