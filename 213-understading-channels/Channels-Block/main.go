package main

import "fmt"

func main() {
	//Channels are the better way to syncronize concurring code

	c := make(chan int)

	c <- 42

	fmt.Println(<-c)

	//channels BLOCK!
}
