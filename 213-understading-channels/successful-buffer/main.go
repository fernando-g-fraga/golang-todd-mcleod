package main

import "fmt"

func main() {
	//Channels are the better way to syncronize concurring code

	c := make(chan int, 1) //Defining C as a channel buffer, it allows a value to be attached on the queue based on the number declared after type.

	c <- 42 //42 is currently queued on the buffer and are ready to fire when it is called

	fmt.Println(<-c)

}