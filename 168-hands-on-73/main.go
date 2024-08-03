package main

import (
	"fmt"
	"time"
)

func main() {

	timedFunction(myapp)

	x := func() {
		fmt.Println("Hello World!")
	}

	Logger(x)

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	callbackfunc := func(num int) {
		fmt.Println("Processing num: ", num)
	}

	ProcessData(data, callbackfunc)

}

func timedFunction(fn func()) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	fmt.Printf("the time elapsed is: %v \n", elapsed)
}

func myapp() {
	time.Sleep(2 * time.Second)
	fmt.Println("Function completed")
}

func Logger(f func()) {
	fmt.Println("Starting execution...")
	f()
	fmt.Println("Execution completed.")
}

func ProcessData(data []int, callback func(int)) {
	for _, item := range data {
		callback(item)
	}
}
