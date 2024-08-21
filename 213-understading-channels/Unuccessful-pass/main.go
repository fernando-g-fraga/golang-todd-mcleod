package main

import "fmt"

func main() {
	//Channels are the better way to syncronize concurring code

	c := make(chan int, 2) //Defining C as a channel buffer, it allows a value to be attached on the queue based on the number declared after type.

	c <- 42 //42 is currently queued on the buffer and are ready to fire when it is called
	c <- 43 //43 is now attempt to be attached on the queue, but we declared the buffer as a size 1 and is already occupied with the value 42

	fmt.Println(<-c)
	fmt.Println(<-c)
}
