package main

import (
	"fmt"
)

func main() {
	//Creating a anonymous function to start a go routine and attach the value on the channel
	c := make(chan int)
	go func() {
		c <- 42
	}()

	//Passing the size of the buffer and attach the value without a func

	/*
		c := make(chan int,1)
			c <- 42
	*/

	fmt.Println(<-c)
}
