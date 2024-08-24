/*
write a program that
○ puts 100 numbers to a channel
○ pull the numbers off the channel and print them
*/

package main

import "fmt"

func main() {
	c := make(chan int)
	receive(c)
	send(c)

}

func receive(c chan int) {
	go func() {
		for i := 0; i < 101; i++ {
			c <- i
		}
		close(c)
	}()
}

func send(c chan int) {
	for i := range c {
		fmt.Println(i)
	}
}
