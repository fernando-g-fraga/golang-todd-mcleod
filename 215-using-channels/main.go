package main

import "fmt"

func main() {

	c := make(chan int)

	go foo(c)

	bar(c)

	fmt.Println("about to exit...")

}

func foo(c chan<- int) { //receiver
	c <- 42
}

func bar(c <-chan int) { //sender
	fmt.Println(<-c)
}
