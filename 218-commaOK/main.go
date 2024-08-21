package main

import "fmt"

func main() {
	odd := make(chan int)
	even := make(chan int)
	quit := make(chan int)

	go receive(odd, even, quit)
	send(odd, even, quit)
}

func receive(odd, even, quit chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(quit)
}

func send(odd, even, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("Receiving from even channel", v)
		case v := <-odd:
			fmt.Println("Receiving from odd channel", v)
		case v, ok := <-quit:
			fmt.Println("Receiving from quit channel", v, ok)
			return
		}
	}
}
