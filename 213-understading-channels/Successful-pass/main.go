package main

import "fmt"

func main() {
	//Channels are the better way to syncronize concurring code

	c := make(chan int)

	//Now the channels are initialize on a go routine that executes on its own, the value 42 is attached to channel and will be availalble to print on func main
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

}
