/*
write a program that
○ launches 10 goroutines
■ each goroutine adds 10 numbers to a channel
○ pull the numbers off the channel and print them
*/
package main

import (
	"fmt"
	"sync"
)

func main() {

	ch1 := make(chan int)
	go receive(ch1)

	for v := range ch1 {
		fmt.Println(v)
	}

	fmt.Println("About to exit!")

}
func receive(ch1 chan int) {
	var wg sync.WaitGroup
	const goRoutines = 10

	wg.Add(goRoutines)
	go func() {
		for i := 0; i < goRoutines; i++ {
			for v := 0; v <= 10; v++ {
				ch1 <- v
			}
			wg.Done()
		}
	}()
	wg.Wait()
	close(ch1)
}
