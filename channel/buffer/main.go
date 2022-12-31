package main

func main()  {
	ch := make(chan string, 2) // argument 2, the channel's buffer size, allows the channel to receive a maximum of two strings
	ch <- "hello"
	ch <- "world" // The application will stop if the buffer size is not more than 1.
	println(<-ch)
	println(<-ch)

}