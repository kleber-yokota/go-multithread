package main

import "fmt"

func main(){
	channel := make(chan string)

	go func(){
		channel <- "Hello world"
	}()

	msg := <- channel
	fmt.Println(msg)
}