package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)
	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("About to Exit!")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- timeToSleep(v2)
			wg.Done()
		}(v)

	}
	wg.Wait()
	close(c2)
}

func timeToSleep(n int) int {

	time.Sleep(time.Microsecond * time.Duration(rand.IntN(500)))
	return n + rand.IntN(1000)
}