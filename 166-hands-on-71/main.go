package main

import "fmt"

func main() {

	y := process(square, 2)

	fmt.Println(y)
}

func process(f func(int) int, n int) string {
	x := f(n)
	return fmt.Sprintf("The number %v squared is %v", n, x)
}

func square(n int) int {
	return n * n
}

/*

Process Function-  \ / - Callback Type Function
					|
					Callback function

*/
