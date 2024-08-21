package main

import (
	"fmt"
	"sync"
)

func main() {

	odd := make(chan int)
	even := make(chan int)
	fanin := make(chan int)

	go send(odd, even)
	go receive(odd, even, fanin)

	for v := range fanin {
		fmt.Println("Value:", v)
	}

}

func send(odd, even chan<- int) {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(odd)
	close(even)
}

func receive(odd, even <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := range odd {
			fanin <- i
		}
		wg.Done()
	}()

	go func() {
		for i := range even {
			fanin <- i
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
