package main

import "fmt"

/*
○ create a func with the identifier foo that returns an int
○ create a func with the identifier bar that returns an int and a string
○ call both funcs
○ print out their results
*/

func main() {

	x := foo(4)
	y, s := bar(5, "Go bears")

	fmt.Println(x)
	fmt.Println(y, s)

}

func foo(n int) int {
	return n
}

func bar(n int, s string) (int, string) {
	return n, s
}
