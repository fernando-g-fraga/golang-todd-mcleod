package main

import "fmt"

func main() {

	odd := make(chan int)
	even := make(chan int)
	quit := make(chan int)

	go send(odd, even, quit)
	receive(odd, even, quit)
	fmt.Println("About to exit!")
}

func send(o, e, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i&2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(o, e, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("The number is even:", v)
		case v := <-o:
			fmt.Println("The number is odd:", v)
		case v := <-q:
			fmt.Println("The number is:", v)
			return
		}
	}
}
