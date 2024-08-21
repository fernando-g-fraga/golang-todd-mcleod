package main

import "fmt"

func main() {

	c := make(chan int)

	//send
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c) //You have to close the channel after you exit the loop to access the receiver of the channel later.
	}()

	//receiver
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("about to exit...")

}
