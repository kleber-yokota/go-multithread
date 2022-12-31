package main

import "fmt"

func receive(name string, hello chan<- string){ //channel just receive data
	hello <- name
}

func read(data <-chan string){ //channel just send data
	fmt.Println(<-data)
}

func main()  {
	hello := make(chan string)
	go receive("Hello", hello)
	read(hello)
}